
#pragma once
class NoCopy {
  public:
    NoCopy(){};
    NoCopy(int val):
    val(val){
    };
    NoCopy(const NoCopy& ) = delete;
    NoCopy( NoCopy&& ) = default;
    int val; 
};