#include <gmock/gmock.h>
#include <vector>
#include <memory>

#include "bar.hpp"
#include "FooMock.hpp"


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
  EXPECT_TRUE(bar.m_Int == 6);
  EXPECT_TRUE(*(bar.m_IntPtr)==7);
  EXPECT_TRUE((bar.m_NoCopyPtr)->val == 8 );
  EXPECT_TRUE( (*(bar.m_NoCopyVector))[0].val == 9 );

  auto vector = bar.GetVector();
  EXPECT_TRUE(vector[0].val==9);
}

int main(int argc, char **argv) {
  ::testing::InitGoogleTest(&argc, argv); 
  return RUN_ALL_TESTS();
}