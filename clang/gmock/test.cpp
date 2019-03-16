#include <gmock/gmock.h>
#include <vector>
#include <memory>


class NoCopy {
  public:
    NoCopy(){};
    NoCopy(int val):
    val(val){
    };
    NoCopy(const NoCopy& ) = delete;
    NoCopy( NoCopy&& ) = default;
    //NoCopy(const NoCopy& iVal) { val = iVal.val;};
    int val; 
};

class Foo 
{
  public:
  virtual ~Foo() = default;
  // virtual void Get(NoCopy& ) = 0 ;
  virtual void GetInt(int& a) = 0;
  virtual int* AllocInt() = 0;
  virtual NoCopy* AllocNoCopy() = 0 ;
  virtual std::vector<NoCopy>* AllocVector() = 0;
  virtual std::vector<NoCopy> BuildVector() = 0;
};

std::vector<NoCopy> kNoCopyCollection;

class FooMock : public Foo{
  public:
  // MOCK_METHOD1(
  //   Get,
  //   void (NoCopy&)
  // );

  std::vector<NoCopy> BuildVector()
  {
    return(*AllocVector());
  }

  MOCK_METHOD1(
    GetInt,
    void (int & )
  );

  MOCK_METHOD0(
    AllocInt,
    int*()
  );

  MOCK_METHOD0(
    AllocNoCopy,
    NoCopy*()
  );

  MOCK_METHOD0(
    AllocVector,
    std::vector<NoCopy>*()
  );
};

class Bar
{
  public:
  Bar(std::shared_ptr<Foo> iFoo):
  foo(iFoo),
  a(2),
  b(3),
  c(new int(2))
  {

  };
  void Do()
  {
    // foo->Get(b);
    foo->GetInt(a);
    c = foo->AllocInt();
    d = foo->AllocNoCopy();
    e = foo->AllocVector();
  };
  std::vector<NoCopy> GetVector()
  {
    return(foo->BuildVector());
  }
  std::shared_ptr<Foo> foo;
  int a;
  NoCopy b;
  int* c;
  NoCopy* d;
  std::vector<NoCopy>* e;
};

NoCopy kNoCopy(5);
int  kInt= 6;
int* kIntptr = new int(7);
NoCopy* kNoCopyptr = new NoCopy(8);
std::vector<NoCopy> * kVector = new std::vector<NoCopy>();
class ExampleTest: public testing::Test
{
public:
  ExampleTest():
  fooMock(std::make_shared<testing::StrictMock<FooMock> >()),
  bar(fooMock)
  {
    NoCopy tmp(9);
    kVector->push_back(std::move(tmp)); 
  };
  std::shared_ptr<FooMock> fooMock;
  Bar bar;


  void expectCall()
  {
    // EXPECT_CALL(
    //   *fooMock,
    //   Get(testing::_)).WillOnce(::testing::SetArgReferee<0>(kNoCopy));

    EXPECT_CALL(
      *fooMock,
      GetInt(::testing::_)).WillOnce(::testing::SetArgReferee<0>(kInt));
    EXPECT_CALL(
      *fooMock,
      AllocInt()
    ).WillOnce(testing::Return(kIntptr));

    EXPECT_CALL(
      *fooMock,
      AllocNoCopy()
    ).WillOnce(testing::Return(kNoCopyptr));

    EXPECT_CALL(
      *fooMock,
      AllocVector()
    ).WillRepeatedly(testing::Return(kVector));
  };

};

TEST_F(ExampleTest, First ) {
  expectCall();
  bar.Do();
  // EXPECT_TRUE(bar.b.val == 5);
  EXPECT_TRUE(bar.a == 6);
  EXPECT_TRUE(*(bar.c)==7);
  EXPECT_TRUE((bar.d)->val == 8 );
  EXPECT_TRUE( (*(bar.e))[0].val == 9 );

  auto vector = bar.GetVector();
  EXPECT_TRUE(vector[0].val==9);
}

int main(int argc, char **argv) {
  ::testing::InitGoogleTest(&argc, argv); 
  return RUN_ALL_TESTS();
}