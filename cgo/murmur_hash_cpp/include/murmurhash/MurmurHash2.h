//-----------------------------------------------------------------------------
// MurmurHash2 was written by Austin Appleby, and is placed in the public
// domain. The author hereby disclaims copyright to this source code.

#ifndef _MURMURHASH2_H_
#define _MURMURHASH2_H_

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
        #ifdef MURMUR_HASH_2_EXPORTS
            #define MURMUR_HASH_2_API extern "C" __declspec(dllexport)
        #else
            #define MURMUR_HASH_2_API extern "C" __declspec(dllimport)
        #endif
    #else
        #ifdef MURMUR_HASH_2_EXPORTS
            #define MURMUR_HASH_2_EXPORTS __declspec(dllexport)
        #else
            #define MURMUR_HASH_2_API __declspec(dllimport)
        #endif
    #endif
#else /* _MSC_VER */
    #ifdef __cplusplus
        #ifdef MURMUR_HASH_2_EXPORTS
            #define MURMUR_HASH_2_API extern "C" __attribute__((visibility ("default")))
        #else
            #define MURMUR_HASH_2_API extern "C"
        #endif
    #else
        #ifdef MURMUR_HASH_2_EXPORTS
            #define MURMUR_HASH_2_API __attribute__((visibility ("default")))
        #else
            #define MURMUR_HASH_2_API
        #endif
    #endif
#endif

//-----------------------------------------------------------------------------

MURMUR_HASH_2_API uint32_t MurmurHash2        ( const void * key, int len, uint32_t seed );
MURMUR_HASH_2_API uint64_t MurmurHash64A      ( const void * key, int len, uint64_t seed );
MURMUR_HASH_2_API uint64_t MurmurHash64B      ( const void * key, int len, uint64_t seed );
MURMUR_HASH_2_API uint32_t MurmurHash2A       ( const void * key, int len, uint32_t seed );
MURMUR_HASH_2_API uint32_t MurmurHashNeutral2 ( const void * key, int len, uint32_t seed );
MURMUR_HASH_2_API uint32_t MurmurHashAligned2 ( const void * key, int len, uint32_t seed );

//-----------------------------------------------------------------------------

#endif // _MURMURHASH2_H_