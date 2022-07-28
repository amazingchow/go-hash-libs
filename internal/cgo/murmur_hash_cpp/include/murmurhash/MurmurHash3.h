//-----------------------------------------------------------------------------
// MurmurHash3 was written by Austin Appleby, and is placed in the public
// domain. The author hereby disclaims copyright to this source code.

#ifndef _MURMURHASH3_H_
#define _MURMURHASH3_H_

//-----------------------------------------------------------------------------
// Platform-specific functions and macros

// Microsoft Visual Studio

#if defined(_MSC_VER) && (_MSC_VER < 1600)

typedef unsigned char uint8_t;
typedef unsigned int uint32_t;
typedef unsigned __int64 uint64_t;

// Other compilers

#else	// defined(_MSC_VER)

#include <stdint.h>

#endif // !defined(_MSC_VER)

#ifdef _MSC_VER
    #ifdef __cplusplus
        #ifdef MURMUR_HASH_3_EXPORTS
            #define MURMUR_HASH_3_API extern "C" __declspec(dllexport)
        #else
            #define MURMUR_HASH_3_API extern "C" __declspec(dllimport)
        #endif
    #else
        #ifdef MURMUR_HASH_3_EXPORTS
            #define MURMUR_HASH_3_EXPORTS __declspec(dllexport)
        #else
            #define MURMUR_HASH_3_API __declspec(dllimport)
        #endif
    #endif
#else /* _MSC_VER */
    #ifdef __cplusplus
        #ifdef MURMUR_HASH_3_EXPORTS
            #define MURMUR_HASH_3_API extern "C" __attribute__((visibility ("default")))
        #else
            #define MURMUR_HASH_3_API extern "C"
        #endif
    #else
        #ifdef MURMUR_HASH_3_EXPORTS
            #define MURMUR_HASH_3_API __attribute__((visibility ("default")))
        #else
            #define MURMUR_HASH_3_API
        #endif
    #endif
#endif

//-----------------------------------------------------------------------------

MURMUR_HASH_3_API void MurmurHash3_x86_32  ( const void * key, int len, uint32_t seed, void * out );

MURMUR_HASH_3_API void MurmurHash3_x86_128 ( const void * key, int len, uint32_t seed, void * out );

MURMUR_HASH_3_API void MurmurHash3_x64_128 ( const void * key, int len, uint32_t seed, void * out );

//-----------------------------------------------------------------------------

#endif // _MURMURHASH3_H_