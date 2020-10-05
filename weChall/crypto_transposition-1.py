## wehChall Training: Crypto - Transposition I (Crypto, Training) 
 
def transposition(ciphertext, key):
        plaintext = ''
        key = list(key)
        numOfBlock = len(ciphertext) / len(key)  # num of block -1
        for block in range(numOfBlock): # block: 0 ~ num of block-1
                for i in key:
                        k = int(i) - 1 + block*len(key)
                        if(k < len(ciphertext)):
                                plaintext += ciphertext[k]
        return plaintext
 
 
 
 
# a is ciphertext and key is "21"
a = "oWdnreuf.lY uoc nar ae dht eemssga eaw yebttrew eh nht eelttre sra enic roertco drre . Ihtni koy uowlu dilekt  oes eoyrup sawsro don:wm cabaocabmg.d"
print "> ciphertext\n"+ a
print "> plaintext\n" + transposition(a, "21")
