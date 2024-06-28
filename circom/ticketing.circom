
pragma circom 2.1.5;

include "circomlib/circuits/comparators.circom";
include "circomlib/circuits/poseidon.circom";
include "circomlib/circuits/eddsaposeidon.circom";

// boolean type field only support hiding or revealing the value.
// The least significant bit represents if the value is hidden or not,
// where 1 means revealed and 0 means hidden.
template BooleanChecker() {
  signal input in;
  signal input hide;
  signal output out;

  component reveal = IsZero();
  reveal.in <== hide;
  out <== in * reveal.out * 2 + reveal.out;
}

// property checker allows caller to prove that input signal's
// equality to a set of values. The result is using the same
// compression schema as id_equals_to:
// The least significant bit is the result of the equality check,
// where 1 means equal and 0 means not equal, and
// the remaining bits are the input signal itself.
template PropertyEqualityChecker(n) {
  signal input in;
  signal input equals_to[n];
  signal output out[n];

  component is_equals[n];
  for (var i = 0; i < n; i++) {
     is_equals[i] = IsEqual();
     is_equals[i].in[0] <== in;
     is_equals[i].in[1] <== equals_to[i];
     out[i] <== is_equals[i].out + equals_to[i] * 2;
  }
}

// ScalarRangeChecker checks if input signal is within the range of [lower_bound, upper_bound], both inclusive.
// There is no output signal. If the input signal is within the range, the circuit will pass, otherwise it will fail.
// NOTE: must do range checks on lower_bound and upper_bound fields
// to make sure that they are within the range of [0, 2**n - 1].
template ScalarRangeChecker(n) {
  signal input in;
  signal input lower_bound;
  signal input upper_bound;

  component lower_bound_checker = GreaterEqThan(n);
  component upper_bound_checker = LessEqThan(n);

  lower_bound_checker.in[0] <== in;
  lower_bound_checker.in[1] <== lower_bound;
  lower_bound_checker.out === 1;
  upper_bound_checker.in[0] <== in;
  upper_bound_checker.in[1] <== upper_bound;
  upper_bound_checker.out === 1;
}

// Scalar256RangeChecker checks if uint256 signal is
// within the range of [lower_bound, upper_bound], both inclusive.
// The uint256 value and bounds are represented as two 128-bit signal.
// NOTE: must do range checks on lower_bound_* and upper_bound_* fields
// to make sure that they are within the range of uint128.
template Scalar256RangeChecker() {
  signal input in_msb;
  signal input in_lsb;
  signal input lower_bound_msb;
  signal input lower_bound_lsb;
  signal input upper_bound_msb;
  signal input upper_bound_lsb;

  component lb_msb_eq = IsEqual();
  component lb_msb_checker = GreaterThan(128);
  component lb_lsb_checker = GreaterEqThan(128);
  // value's msb is greater or equal than lower_bound's msb
  lb_msb_checker.in[0] <== in_msb;
  lb_msb_checker.in[1] <== lower_bound_msb;
  lb_msb_eq.in[0] <== in_msb;
  lb_msb_eq.in[1] <== lower_bound_msb;
  // value's lsb is greater or equal than lower_bound's lsb
  lb_lsb_checker.in[0] <== in_lsb;
  lb_lsb_checker.in[1] <== lower_bound_lsb;
  // either value's msb is greater or msb is equal and lsb is greater or equal
  lb_msb_checker.out + lb_msb_eq.out * lb_lsb_checker.out === 1;

  component up_msb_eq = IsEqual();
  component up_msb_checker = LessThan(128);
  component up_lsb_checker = LessEqThan(128);
  // value's msb is less or equal than upper_bound's msb
  up_msb_checker.in[0] <== in_msb;
  up_msb_checker.in[1] <== upper_bound_msb;
  up_msb_eq.in[0] <== in_msb;
  up_msb_eq.in[1] <== upper_bound_msb;
  // value's lsb is less or equal than upper_bound's lsb
  up_lsb_checker.in[0] <== in_lsb;
  up_lsb_checker.in[1] <== upper_bound_lsb;
  // either value's msb is less or is equal and lsb is less or equal
  up_msb_checker.out + up_msb_eq.out * up_lsb_checker.out === 1;
}

template AllMetadataHasher() {
  signal input version;
  signal input type;
  signal input context;
  signal input id;
  signal input verification_stack;
  signal input signature_id;
  signal input expired_at;
  signal input identity_commitment;

  signal output out;

  component hasher = Poseidon(8);
  hasher.inputs[0] <== version;
  hasher.inputs[1] <== type;
  hasher.inputs[2] <== context;
  hasher.inputs[3] <== id;
  hasher.inputs[4] <== verification_stack;
  hasher.inputs[5] <== signature_id;
  hasher.inputs[6] <== expired_at;
  hasher.inputs[7] <== identity_commitment;
  out <== hasher.out;
}

template CredHasher() {
  signal input metadata_hash;
  signal input body_hash;

  signal output out;

  component hasher = Poseidon(2);
  hasher.inputs[0] <== metadata_hash;
  hasher.inputs[1] <== body_hash;
  out <== hasher.out;
}

template BodyHasher() {
  signal output out;

  out <== 0;
}

template Main() {
  // headers
  signal input ticket_version;
  signal input ticket_type;
  signal input ticket_context;
  signal input ticket_id;

  // signature metadata
  signal input ticket_sig_verification_stack;
  signal input ticket_sig_id;
  signal input ticket_sig_expired_at;
  signal input ticket_sig_identity_commitment;

  // signature
  signal input ticket_sig_pubkey_x;
  signal input ticket_sig_pubkey_y;
  signal input ticket_sig_s;
  signal input ticket_sig_r8_x;
  signal input ticket_sig_r8_y;

  // verification input
  signal input ticket_identity_secret;
  signal input ticket_internal_nullifier;
  signal input ticket_external_nullifier;
  
  // identity result
  signal input ticket_revealing_identity;

  // primitive control signals
  signal input ticket_expiration_lb;
  signal input ticket_id_equals_to;

  // email credential inputs
  // headers
  signal input email_version;
  signal input email_type;
  signal input email_context;
  signal input email_id;

  // signature metadata
  signal input email_sig_verification_stack;
  signal input email_sig_id;
  signal input email_sig_expired_at;
  signal input email_sig_identity_commitment;

  // signature
  signal input email_sig_pubkey_x;
  signal input email_sig_pubkey_y;
  signal input email_sig_s;
  signal input email_sig_r8_x;
  signal input email_sig_r8_y;

  // verification input
  signal input email_identity_secret;
  signal input email_internal_nullifier;
  signal input email_external_nullifier;
  
  // identity result
  signal input email_revealing_identity;
  // HMAC from poseidon(identity_secret, external_nullifier, revealing_identity)
  signal input email_revealing_identity_hmac;

  // primitive control signals
  signal input email_expiration_lb;

  // intrinsic output signals
  signal output ticket_out_context;
  signal output ticket_out_expiration_lb;
  signal output ticket_out_key_id;

  signal output email_out_context;
  signal output email_out_nullifier;
  signal output email_out_external_nullifier;
  signal output email_out_revealing_identity;
  signal output email_out_expiration_lb;
  signal output email_out_key_id;

  // defs

  // basic checks
  ticket_version === 1;  // protocol version 1
  ticket_sig_verification_stack === 1;  // babyzk
  ticket_type === 1;

  ticket_sig_identity_commitment === 0; // ticket credential does not have identity commitment

  email_version === 1;  // protocol version 1
  email_sig_verification_stack === 1;  // babyzk
  email_type === 1;

  // email address check
  email_id === ticket_id;

  // redirect intrinsic signals to output
  ticket_out_context <== ticket_context;
  ticket_out_expiration_lb <== ticket_expiration_lb;

  email_out_context <== email_context;
  email_out_external_nullifier <== email_external_nullifier;
  email_out_revealing_identity <== email_revealing_identity;
  email_out_expiration_lb <== email_expiration_lb;

  component ticket_all_metadata_hasher = AllMetadataHasher();
  ticket_all_metadata_hasher.version <== ticket_version;
  ticket_all_metadata_hasher.type <== ticket_type;
  ticket_all_metadata_hasher.context <== ticket_context;
  ticket_all_metadata_hasher.id <== ticket_id;
  ticket_all_metadata_hasher.verification_stack <== ticket_sig_verification_stack;
  ticket_all_metadata_hasher.signature_id <== ticket_sig_id;
  ticket_all_metadata_hasher.expired_at <== ticket_sig_expired_at;
  ticket_all_metadata_hasher.identity_commitment <== ticket_sig_identity_commitment;

  component email_all_metadata_hasher = AllMetadataHasher();
  email_all_metadata_hasher.version <== email_version;
  email_all_metadata_hasher.type <== email_type;
  email_all_metadata_hasher.context <== email_context;
  email_all_metadata_hasher.id <== email_id;
  email_all_metadata_hasher.verification_stack <== email_sig_verification_stack;
  email_all_metadata_hasher.signature_id <== email_sig_id;
  email_all_metadata_hasher.expired_at <== email_sig_expired_at;
  email_all_metadata_hasher.identity_commitment <== email_sig_identity_commitment;

  component body_hasher = BodyHasher();

  component ticket_cred_hasher = CredHasher();
  ticket_cred_hasher.metadata_hash <== ticket_all_metadata_hasher.out;
  ticket_cred_hasher.body_hash <== body_hasher.out;

  component email_cred_hasher = CredHasher();
  email_cred_hasher.metadata_hash <== email_all_metadata_hasher.out;
  email_cred_hasher.body_hash <== body_hasher.out;

  component ticket_eddsa = EdDSAPoseidonVerifier();
  ticket_eddsa.enabled <== 1;
  ticket_eddsa.M <== ticket_cred_hasher.out;
  ticket_eddsa.Ax <== ticket_sig_pubkey_x;
  ticket_eddsa.Ay <== ticket_sig_pubkey_y;
  ticket_eddsa.R8x <== ticket_sig_r8_x;
  ticket_eddsa.R8y <== ticket_sig_r8_y;
  ticket_eddsa.S <== ticket_sig_s;

  component email_eddsa = EdDSAPoseidonVerifier();
  email_eddsa.enabled <== 1;
  email_eddsa.M <== email_cred_hasher.out;
  email_eddsa.Ax <== email_sig_pubkey_x;
  email_eddsa.Ay <== email_sig_pubkey_y;
  email_eddsa.R8x <== email_sig_r8_x;
  email_eddsa.R8y <== email_sig_r8_y;
  email_eddsa.S <== email_sig_s;

  // verification output
  component ticket_key_id_hasher = Poseidon(2);
  ticket_key_id_hasher.inputs[0] <== ticket_sig_pubkey_x;
  ticket_key_id_hasher.inputs[1] <== ticket_sig_pubkey_y;
  ticket_out_key_id <== ticket_key_id_hasher.out;

  component email_key_id_hasher = Poseidon(2);
  email_key_id_hasher.inputs[0] <== email_sig_pubkey_x;
  email_key_id_hasher.inputs[1] <== email_sig_pubkey_y;
  email_out_key_id <== email_key_id_hasher.out;

  // primitive controls
  component ticket_expiration_gte = GreaterEqThan(64);
  ticket_expiration_gte.in[0] <== ticket_sig_expired_at;
  ticket_expiration_gte.in[1] <== ticket_expiration_lb;
  ticket_expiration_gte.out === 1;

  component email_expiration_gte = GreaterEqThan(64);
  email_expiration_gte.in[0] <== email_sig_expired_at;
  email_expiration_gte.in[1] <== email_expiration_lb;
  email_expiration_gte.out === 1;

  // skip ticket credential identity commitment check
  component is_id_valid = Poseidon(2);
  is_id_valid.inputs[0] <== email_identity_secret;
  is_id_valid.inputs[1] <== email_internal_nullifier;
  is_id_valid.out === email_sig_identity_commitment;
  
  component email_nullifier_hasher = Poseidon(2);
  email_nullifier_hasher.inputs[0] <== email_internal_nullifier;
  email_nullifier_hasher.inputs[1] <== email_external_nullifier;
  email_out_nullifier <== email_nullifier_hasher.out;
  
  component email_revealing_identity_hmac_check = Poseidon(3);
  email_revealing_identity_hmac_check.inputs[0] <== email_identity_secret;
  email_revealing_identity_hmac_check.inputs[1] <== email_external_nullifier;
  email_revealing_identity_hmac_check.inputs[2] <== email_revealing_identity;
  email_revealing_identity_hmac_check.out === email_revealing_identity_hmac;
}

component main = Main();
