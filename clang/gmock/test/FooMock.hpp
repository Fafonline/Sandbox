#pragma once
#include <gmock/gmock.h>
#include "Foo.hpp"
class FooMock : public Foo{
  public:
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