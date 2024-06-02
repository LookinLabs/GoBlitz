import os
import binascii

def generate_api_key():
    return binascii.hexlify(os.urandom(16)).decode()

print(generate_api_key())