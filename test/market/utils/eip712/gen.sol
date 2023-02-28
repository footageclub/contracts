	function _computeMerkleProofDepth7(Eip712MerkleProof proofPtr, uint256 leaf) pure returns (bytes32 root) {
		assembly{
		let key := shr(248, proofPtr)
		let proof := add(proofPtr, 1)
			// Compute level 1
			let scratch := shl(5, and(key, 1))
			mstore(scratch, leaf)
			mstore(xor(scratch, OneWord), calldataload(proof))

			// Compute level 2
			scratch := shl(5, and(shr(1, key), 1))
			mstore(scratch, keccak256(0, TwoWords))
			mstore(xor(scratch, OneWord), calldataload(add(proof, 0x20)))

			// Compute level 3
			scratch := shl(5, and(shr(2, key), 1))
			mstore(scratch, keccak256(0, TwoWords))
			mstore(xor(scratch, OneWord), calldataload(add(proof, 0x40)))

			// Compute level 4
			scratch := shl(5, and(shr(3, key), 1))
			mstore(scratch, keccak256(0, TwoWords))
			mstore(xor(scratch, OneWord), calldataload(add(proof, 0x60)))

			// Compute level 5
			scratch := shl(5, and(shr(4, key), 1))
			mstore(scratch, keccak256(0, TwoWords))
			mstore(xor(scratch, OneWord), calldataload(add(proof, 0x80)))

			// Compute level 6
			scratch := shl(5, and(shr(5, key), 1))
			mstore(scratch, keccak256(0, TwoWords))
			mstore(xor(scratch, OneWord), calldataload(add(proof, 0xa0)))

			// Compute root hash
			scratch := shl(5, and(shr(6, key), 1))
			mstore(scratch, keccak256(0, TwoWords))
			mstore(xor(scratch, OneWord), calldataload(add(proof, 0xc0)))
			root := keccak256(0, TwoWords)
		}
	}