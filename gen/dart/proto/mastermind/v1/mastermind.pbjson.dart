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
@$core.Deprecated('Use agilityLevelDescriptor instead')
const AgilityLevel$json = const {
  '1': 'AgilityLevel',
  '2': const [
    const {'1': 'AGILITY_LEVEL_UNSPECIFIED', '2': 0},
    const {'1': 'AGILITY_LEVEL_HIGH', '2': 1},
    const {'1': 'AGILITY_LEVEL_MEDIUM', '2': 2},
    const {'1': 'AGILITY_LEVEL_LOW', '2': 3},
  ],
};

/// Descriptor for `AgilityLevel`. Decode as a `google.protobuf.EnumDescriptorProto`.
final $typed_data.Uint8List agilityLevelDescriptor = $convert.base64Decode('CgxBZ2lsaXR5TGV2ZWwSHQoZQUdJTElUWV9MRVZFTF9VTlNQRUNJRklFRBAAEhYKEkFHSUxJVFlfTEVWRUxfSElHSBABEhgKFEFHSUxJVFlfTEVWRUxfTUVESVVNEAISFQoRQUdJTElUWV9MRVZFTF9MT1cQAw==');
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
@$core.Deprecated('Use planeDescriptor instead')
const Plane$json = const {
  '1': 'Plane',
  '2': const [
    const {'1': 'plane_id', '3': 1, '4': 1, '5': 5, '10': 'planeId'},
    const {'1': 'lat', '3': 2, '4': 1, '5': 1, '10': 'lat'},
    const {'1': 'lon', '3': 3, '4': 1, '5': 1, '10': 'lon'},
    const {'1': 'alt', '3': 4, '4': 1, '5': 1, '10': 'alt'},
    const {'1': 'heading', '3': 5, '4': 1, '5': 1, '10': 'heading'},
    const {'1': 'speed', '3': 6, '4': 1, '5': 1, '10': 'speed'},
    const {'1': 'agility_level', '3': 7, '4': 1, '5': 14, '6': '.mastermind.v1.AgilityLevel', '10': 'agilityLevel'},
    const {'1': 'score', '3': 8, '4': 1, '5': 1, '10': 'score'},
  ],
};

/// Descriptor for `Plane`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List planeDescriptor = $convert.base64Decode('CgVQbGFuZRIZCghwbGFuZV9pZBgBIAEoBVIHcGxhbmVJZBIQCgNsYXQYAiABKAFSA2xhdBIQCgNsb24YAyABKAFSA2xvbhIQCgNhbHQYBCABKAFSA2FsdBIYCgdoZWFkaW5nGAUgASgBUgdoZWFkaW5nEhQKBXNwZWVkGAYgASgBUgVzcGVlZBJACg1hZ2lsaXR5X2xldmVsGAcgASgOMhsubWFzdGVybWluZC52MS5BZ2lsaXR5TGV2ZWxSDGFnaWxpdHlMZXZlbBIUCgVzY29yZRgIIAEoAVIFc2NvcmU=');
@$core.Deprecated('Use getPlanesRequestDescriptor instead')
const GetPlanesRequest$json = const {
  '1': 'GetPlanesRequest',
};

/// Descriptor for `GetPlanesRequest`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List getPlanesRequestDescriptor = $convert.base64Decode('ChBHZXRQbGFuZXNSZXF1ZXN0');
@$core.Deprecated('Use getPlanesResponseDescriptor instead')
const GetPlanesResponse$json = const {
  '1': 'GetPlanesResponse',
  '2': const [
    const {'1': 'selected_plane_id', '3': 1, '4': 1, '5': 5, '10': 'selectedPlaneId'},
    const {'1': 'planes', '3': 2, '4': 3, '5': 11, '6': '.mastermind.v1.Plane', '10': 'planes'},
  ],
};

/// Descriptor for `GetPlanesResponse`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List getPlanesResponseDescriptor = $convert.base64Decode('ChFHZXRQbGFuZXNSZXNwb25zZRIqChFzZWxlY3RlZF9wbGFuZV9pZBgBIAEoBVIPc2VsZWN0ZWRQbGFuZUlkEiwKBnBsYW5lcxgCIAMoCzIULm1hc3Rlcm1pbmQudjEuUGxhbmVSBnBsYW5lcw==');
