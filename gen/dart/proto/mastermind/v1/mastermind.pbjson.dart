///
//  Generated code. Do not modify.
//  source: proto/mastermind/v1/mastermind.proto
//
// @dart = 2.12
// ignore_for_file: annotate_overrides,camel_case_types,constant_identifier_names,deprecated_member_use_from_same_package,directives_ordering,library_prefixes,non_constant_identifier_names,prefer_final_fields,return_of_invalid_type,unnecessary_const,unnecessary_import,unnecessary_this,unused_import,unused_shown_name

import 'dart:core' as $core;
import 'dart:convert' as $convert;
import 'dart:typed_data' as $typed_data;
@$core.Deprecated('Use missionStateDescriptor instead')
const MissionState$json = const {
  '1': 'MissionState',
  '2': const [
    const {'1': 'MISSION_STATE_UNSPECIFIED', '2': 0},
    const {'1': 'MISSION_STATE_APPROACH', '2': 1},
    const {'1': 'MISSION_STATE_FOLLOWING', '2': 2},
    const {'1': 'MISSION_STATE_KAMIKAZE', '2': 3},
  ],
};

/// Descriptor for `MissionState`. Decode as a `google.protobuf.EnumDescriptorProto`.
final $typed_data.Uint8List missionStateDescriptor = $convert.base64Decode('CgxNaXNzaW9uU3RhdGUSHQoZTUlTU0lPTl9TVEFURV9VTlNQRUNJRklFRBAAEhoKFk1JU1NJT05fU1RBVEVfQVBQUk9BQ0gQARIbChdNSVNTSU9OX1NUQVRFX0ZPTExPV0lORxACEhoKFk1JU1NJT05fU1RBVEVfS0FNSUtBWkUQAw==');
@$core.Deprecated('Use stateTransitionDescriptor instead')
const StateTransition$json = const {
  '1': 'StateTransition',
  '2': const [
    const {'1': 'STATE_TRANSITION_UNSPECIFIED', '2': 0},
    const {'1': 'STATE_TRANSITION_TARGET_APPROACHED', '2': 1},
    const {'1': 'STATE_TRANSITION_LOCK_FAILED', '2': 2},
    const {'1': 'STATE_TRANSITION_LOCK_SUCCESS', '2': 3},
    const {'1': 'STATE_TRANSITION_QR_SUCCESS', '2': 4},
    const {'1': 'STATE_TRANSITION_QR_FAILED', '2': 5},
    const {'1': 'STATE_TRANSITION_MODE_KAMIKAZE_SELECTED', '2': 6},
    const {'1': 'STATE_TRANSITION_MODE_APPROACH_SELECTED', '2': 7},
  ],
};

/// Descriptor for `StateTransition`. Decode as a `google.protobuf.EnumDescriptorProto`.
final $typed_data.Uint8List stateTransitionDescriptor = $convert.base64Decode('Cg9TdGF0ZVRyYW5zaXRpb24SIAocU1RBVEVfVFJBTlNJVElPTl9VTlNQRUNJRklFRBAAEiYKIlNUQVRFX1RSQU5TSVRJT05fVEFSR0VUX0FQUFJPQUNIRUQQARIgChxTVEFURV9UUkFOU0lUSU9OX0xPQ0tfRkFJTEVEEAISIQodU1RBVEVfVFJBTlNJVElPTl9MT0NLX1NVQ0NFU1MQAxIfChtTVEFURV9UUkFOU0lUSU9OX1FSX1NVQ0NFU1MQBBIeChpTVEFURV9UUkFOU0lUSU9OX1FSX0ZBSUxFRBAFEisKJ1NUQVRFX1RSQU5TSVRJT05fTU9ERV9LQU1JS0FaRV9TRUxFQ1RFRBAGEisKJ1NUQVRFX1RSQU5TSVRJT05fTU9ERV9BUFBST0FDSF9TRUxFQ1RFRBAH');
@$core.Deprecated('Use updateStateRequestDescriptor instead')
const UpdateStateRequest$json = const {
  '1': 'UpdateStateRequest',
  '2': const [
    const {'1': 'state_transition', '3': 1, '4': 1, '5': 14, '6': '.mastermind.v1.StateTransition', '10': 'stateTransition'},
  ],
};

/// Descriptor for `UpdateStateRequest`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List updateStateRequestDescriptor = $convert.base64Decode('ChJVcGRhdGVTdGF0ZVJlcXVlc3QSSQoQc3RhdGVfdHJhbnNpdGlvbhgBIAEoDjIeLm1hc3Rlcm1pbmQudjEuU3RhdGVUcmFuc2l0aW9uUg9zdGF0ZVRyYW5zaXRpb24=');
@$core.Deprecated('Use updateStateResponseDescriptor instead')
const UpdateStateResponse$json = const {
  '1': 'UpdateStateResponse',
  '2': const [
    const {'1': 'old_state', '3': 1, '4': 1, '5': 14, '6': '.mastermind.v1.MissionState', '10': 'oldState'},
    const {'1': 'state_transition', '3': 2, '4': 1, '5': 14, '6': '.mastermind.v1.StateTransition', '10': 'stateTransition'},
    const {'1': 'current_state', '3': 3, '4': 1, '5': 14, '6': '.mastermind.v1.MissionState', '10': 'currentState'},
  ],
};

/// Descriptor for `UpdateStateResponse`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List updateStateResponseDescriptor = $convert.base64Decode('ChNVcGRhdGVTdGF0ZVJlc3BvbnNlEjgKCW9sZF9zdGF0ZRgBIAEoDjIbLm1hc3Rlcm1pbmQudjEuTWlzc2lvblN0YXRlUghvbGRTdGF0ZRJJChBzdGF0ZV90cmFuc2l0aW9uGAIgASgOMh4ubWFzdGVybWluZC52MS5TdGF0ZVRyYW5zaXRpb25SD3N0YXRlVHJhbnNpdGlvbhJACg1jdXJyZW50X3N0YXRlGAMgASgOMhsubWFzdGVybWluZC52MS5NaXNzaW9uU3RhdGVSDGN1cnJlbnRTdGF0ZQ==');
const $core.Map<$core.String, $core.dynamic> MastermindServiceBase$json = const {
  '1': 'MastermindService',
  '2': const [
    const {'1': 'UpdateState', '2': '.mastermind.v1.UpdateStateRequest', '3': '.mastermind.v1.UpdateStateResponse', '5': true, '6': true},
  ],
};

@$core.Deprecated('Use mastermindServiceDescriptor instead')
const $core.Map<$core.String, $core.Map<$core.String, $core.dynamic>> MastermindServiceBase$messageJson = const {
  '.mastermind.v1.UpdateStateRequest': UpdateStateRequest$json,
  '.mastermind.v1.UpdateStateResponse': UpdateStateResponse$json,
};

/// Descriptor for `MastermindService`. Decode as a `google.protobuf.ServiceDescriptorProto`.
final $typed_data.Uint8List mastermindServiceDescriptor = $convert.base64Decode('ChFNYXN0ZXJtaW5kU2VydmljZRJYCgtVcGRhdGVTdGF0ZRIhLm1hc3Rlcm1pbmQudjEuVXBkYXRlU3RhdGVSZXF1ZXN0GiIubWFzdGVybWluZC52MS5VcGRhdGVTdGF0ZVJlc3BvbnNlKAEwAQ==');
