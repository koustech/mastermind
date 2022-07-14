///
//  Generated code. Do not modify.
//  source: mastermind/v1/mastermind.proto
//
// @dart = 2.12
// ignore_for_file: annotate_overrides,camel_case_types,constant_identifier_names,directives_ordering,library_prefixes,non_constant_identifier_names,prefer_final_fields,return_of_invalid_type,unnecessary_const,unnecessary_import,unnecessary_this,unused_import,unused_shown_name

// ignore_for_file: UNDEFINED_SHOWN_NAME
import 'dart:core' as $core;
import 'package:protobuf/protobuf.dart' as $pb;

class MissionState extends $pb.ProtobufEnum {
  static const MissionState MISSION_STATE_UNSPECIFIED = MissionState._(0, const $core.bool.fromEnvironment('protobuf.omit_enum_names') ? '' : 'MISSION_STATE_UNSPECIFIED');
  static const MissionState MISSION_STATE_APPROACH = MissionState._(1, const $core.bool.fromEnvironment('protobuf.omit_enum_names') ? '' : 'MISSION_STATE_APPROACH');
  static const MissionState MISSION_STATE_FOLLOWING = MissionState._(2, const $core.bool.fromEnvironment('protobuf.omit_enum_names') ? '' : 'MISSION_STATE_FOLLOWING');
  static const MissionState MISSION_STATE_KAMIKAZE = MissionState._(3, const $core.bool.fromEnvironment('protobuf.omit_enum_names') ? '' : 'MISSION_STATE_KAMIKAZE');

  static const $core.List<MissionState> values = <MissionState> [
    MISSION_STATE_UNSPECIFIED,
    MISSION_STATE_APPROACH,
    MISSION_STATE_FOLLOWING,
    MISSION_STATE_KAMIKAZE,
  ];

  static final $core.Map<$core.int, MissionState> _byValue = $pb.ProtobufEnum.initByValue(values);
  static MissionState? valueOf($core.int value) => _byValue[value];

  const MissionState._($core.int v, $core.String n) : super(v, n);
}

class StateTransition extends $pb.ProtobufEnum {
  static const StateTransition STATE_TRANSITION_UNSPECIFIED = StateTransition._(0, const $core.bool.fromEnvironment('protobuf.omit_enum_names') ? '' : 'STATE_TRANSITION_UNSPECIFIED');
  static const StateTransition STATE_TRANSITION_TARGET_APPROACHED = StateTransition._(1, const $core.bool.fromEnvironment('protobuf.omit_enum_names') ? '' : 'STATE_TRANSITION_TARGET_APPROACHED');
  static const StateTransition STATE_TRANSITION_LOCK_FAILED = StateTransition._(2, const $core.bool.fromEnvironment('protobuf.omit_enum_names') ? '' : 'STATE_TRANSITION_LOCK_FAILED');
  static const StateTransition STATE_TRANSITION_LOCK_SUCCESS = StateTransition._(3, const $core.bool.fromEnvironment('protobuf.omit_enum_names') ? '' : 'STATE_TRANSITION_LOCK_SUCCESS');
  static const StateTransition STATE_TRANSITION_QR_SUCCESS = StateTransition._(4, const $core.bool.fromEnvironment('protobuf.omit_enum_names') ? '' : 'STATE_TRANSITION_QR_SUCCESS');
  static const StateTransition STATE_TRANSITION_QR_FAILED = StateTransition._(5, const $core.bool.fromEnvironment('protobuf.omit_enum_names') ? '' : 'STATE_TRANSITION_QR_FAILED');
  static const StateTransition STATE_TRANSITION_MODE_KAMIKAZE_SELECTED = StateTransition._(6, const $core.bool.fromEnvironment('protobuf.omit_enum_names') ? '' : 'STATE_TRANSITION_MODE_KAMIKAZE_SELECTED');
  static const StateTransition STATE_TRANSITION_MODE_APPROACH_SELECTED = StateTransition._(7, const $core.bool.fromEnvironment('protobuf.omit_enum_names') ? '' : 'STATE_TRANSITION_MODE_APPROACH_SELECTED');

  static const $core.List<StateTransition> values = <StateTransition> [
    STATE_TRANSITION_UNSPECIFIED,
    STATE_TRANSITION_TARGET_APPROACHED,
    STATE_TRANSITION_LOCK_FAILED,
    STATE_TRANSITION_LOCK_SUCCESS,
    STATE_TRANSITION_QR_SUCCESS,
    STATE_TRANSITION_QR_FAILED,
    STATE_TRANSITION_MODE_KAMIKAZE_SELECTED,
    STATE_TRANSITION_MODE_APPROACH_SELECTED,
  ];

  static final $core.Map<$core.int, StateTransition> _byValue = $pb.ProtobufEnum.initByValue(values);
  static StateTransition? valueOf($core.int value) => _byValue[value];

  const StateTransition._($core.int v, $core.String n) : super(v, n);
}

