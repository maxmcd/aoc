import peachpy
from peachpy.x86_64 import *

# Function to detect if a string contains unique elements
@peachpy.function
def unique_elements(string: str, length: int):
    # Set up function arguments
    string_pointer = Argument(ptr(const_char))
    length_pointer = Argument(ptr(const_int))

    # Load arguments
    with Function("unique_elements", (string_pointer, length_pointer)):
        # Load the string and length
        string_pointer = GeneralPurposeRegister.rdi
        length_pointer = GeneralPurposeRegister.rsi
        string = mem_ptr[string_pointer]
        length = mem_ptr[length_pointer]

        # Set up loop counters and a bitmap to track seen characters
        i = GeneralPurposeRegister.r8
        j = GeneralPurposeRegister.r9
        bitmap = GeneralPurposeRegister.r10
        mov(bitmap, 0)

        # Loop through the string, checking each character against the bitmap
        for char in string:
            # Shift the bitmap and set the bit for the current character
            shl(bitmap, 1)
            or_(bitmap, 1 << (ord(char) % 64))

            # Check if the bitmap already has the bit set for this character
            jz(".not_unique")

        # If we reach the end of the loop, the string has unique elements
        mov(al, 1)
        ret()

        # If we encounter a non-unique character, return 0
        LABEL(".not_unique")
        mov(al, 0)
        ret()
