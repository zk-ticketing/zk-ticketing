# compilation
compile: ticketing.circom
	circom ticketing.circom --r1cs --wasm --sym --c -o build
print-r1cs:
	snarkjs r1cs info build/ticketing.r1cs
print-r1cs-json:
	snarkjs r1cs export json build/tlicketing.r1cs build/ticketing.r1cs.json && cat build/ticketing.r1cs.json
print-sym:
	snarkjs r1cs print build/ticketing.r1cs build/ticketing.sym
gen-witness:
	cd build/ticketing_js && node generate_witness.js ticketing.wasm ../../test/example.json ../witness.wtns
	# cd build && snarkjs wtns check ticketing.r1cs witness.wtns

# phase2 setup for groth16
# powersOfTau28_hez_final_14 - supports up to 16k constraint
init-groth16-16k: gen-witness
	snarkjs groth16 setup build/ticketing.r1cs ptaus/powersOfTau28_hez_final_14.ptau build/ticketing_0000.zkey
contribute-groth16:
	snarkjs zkey contribute build/ticketing_0000.zkey build/ticketing_0001.zkey --name="p2 1st Contributor Name" -v
finalize-groth16:
	snarkjs zkey beacon build/ticketing_0001.zkey build/ticketing_final.zkey 0102030405060708090a0b0c0d0e0f101112131415161718191a1b1c1d1e1f 10 -n="Final Beacon phase2"
do-export-zkey:
	snarkjs zkey export verificationkey build/ticketing_final.zkey build/verification_key.json
whole-phase2:
	make init-groth16-16k
	make contribute-groth16
	make finalize-groth16
	make do-export-zkey
verify-final-zkey-16k:
	snarkjs zkey verify build/ticketing.r1cs ptaus/powersOfTau28_hez_final_14.ptau build/ticketing_final.zkey

# proof generation and verification
gen-proof:
	snarkjs groth16 prove build/ticketing_final.zkey build/witness.wtns build/proof.json build/public.json
verify-proof:
	snarkjs groth16 verify build/verification_key.json build/public.json build/proof.json
gen-full-proof:
	snarkjs groth16 fullprove passport_example.json build/ticketing_js/ticketing.wasm build/ticketing_final.zkey build/proof.json build/public.json
gen-sol:
	snarkjs zkey export solidityverifier build/ticketing_final.zkey build/verifier.sol
