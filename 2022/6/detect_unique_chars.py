import peachpy
import peachpy.x86_64

# Define a function that takes a pointer to a string and its length
@peachpy.x86_64.function('detect_unique_chars')
def detect_unique_chars(string_ptr: peachpy.x86_64.abi.pointer, length: peachpy.x86_64.abi.size_t) -> peachpy.x86_64.abi.int:
    # Load the length of the string into a register
    length_reg = peachpy.x86_64.GeneralPurposeRegister64()
    peachpy.x86_64.MOV(length_reg, length)

    # Load the pointer to the string into a register
    string_ptr_reg = peachpy.x86_64.GeneralPurposeRegister64()
    peachpy.x86_64.MOV(string_ptr_reg, string_ptr)

    # Initialize a register to store the result
    result_reg = peachpy.x86_64.GeneralPurposeRegister64()
    peachpy.x86_64.MOV(result_reg, 0)

    # Loop through the string by incrementing the pointer and decrementing the length
    loop_label = peachpy.x86_64.Label()
    peachpy.x86_64.LABEL(loop_label)
    with peachpy.x86_64.If(length_reg > 0):
        # Load the current character into a register
        current_char_reg = peachpy.x86_64.AVXRegister(1)
        peachpy.x86_64.VMOV(current_char_reg, [string_ptr_reg])

        # XOR the current character with the result to "add" it to the result
        peachpy.x86_64.VPXOR(current_char_reg, result_reg, current_char_reg)

        # Store the result back in the result register
        peachpy.x86_64.VMOV(result_reg, current_char_reg)

        # Increment the pointer and decrement the length
        peachpy.x86_64.ADD(string_ptr_reg, 1)
        peachpy.x86_64.DEC(length_reg)

        # Jump back to the beginning of the loop
        peachpy.x86_64.JMP(loop_label)

    # Return the result
    peachpy.x86_64.RETURN(result_reg)
