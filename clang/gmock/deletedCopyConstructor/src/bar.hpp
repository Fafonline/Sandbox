#pragma once

#include <memory>
#include "Foo.hpp"

class Bar
{
  public:
  Bar(std::shared_ptr<Foo> iFoo):
  foo(iFoo),
  m_Int(2),
  m_NoCopy(3),
  m_IntPtr(new int(2))
  {

  };
  void Do()
  {
    // foo->Get(b);
    foo->GetInt(m_Int);
    m_IntPtr = foo->AllocInt();
    m_NoCopyPtr = foo->AllocNoCopy();
    m_NoCopyVector = foo->AllocVector();
  };
  std::vector<NoCopy> GetVector()
  {
    return(foo->BuildVector());
  }
  std::shared_ptr<Foo> foo;
  int m_Int;
  NoCopy m_NoCopy;
  int* m_IntPtr;
  NoCopy* m_NoCopyPtr;
  std::vector<NoCopy>* m_NoCopyVector;
};
