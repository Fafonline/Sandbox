#pragma once
#include "NoCopy.hpp"
#include <vector>
class Foo 
{
  public:
  virtual ~Foo() = default;
  virtual void GetInt(int& a) = 0;
  virtual int* AllocInt() = 0;
  virtual NoCopy* AllocNoCopy() = 0 ;
  virtual std::vector<NoCopy>* AllocVector() = 0;
  virtual std::vector<NoCopy> BuildVector() = 0;
};