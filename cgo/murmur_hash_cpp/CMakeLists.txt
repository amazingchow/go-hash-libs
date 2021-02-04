cmake_minimum_required(VERSION 3.5)

project(murmur_hash_cpp)

add_library(murmurhash2 SHARED 
    src/MurmurHash2.cpp
)

target_include_directories(murmurhash2
    PUBLIC 
        ${PROJECT_SOURCE_DIR}/include
)

add_library(murmurhash3 SHARED 
    src/MurmurHash3.cpp
)

target_include_directories(murmurhash3
    PUBLIC 
        ${PROJECT_SOURCE_DIR}/include
)

install (TARGETS murmurhash2
    LIBRARY DESTINATION lib)

install (TARGETS murmurhash3
    LIBRARY DESTINATION lib)

install(DIRECTORY ${PROJECT_SOURCE_DIR}/include/ 
    DESTINATION include)
