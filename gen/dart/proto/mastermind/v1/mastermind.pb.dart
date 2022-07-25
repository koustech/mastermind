///
//  Generated code. Do not modify.
//  source: proto/mastermind/v1/mastermind.proto
//
// @dart = 2.12
// ignore_for_file: annotate_overrides,camel_case_types,constant_identifier_names,directives_ordering,library_prefixes,non_constant_identifier_names,prefer_final_fields,return_of_invalid_type,unnecessary_const,unnecessary_import,unnecessary_this,unused_import,unused_shown_name

import 'dart:core' as $core;

import 'package:protobuf/protobuf.dart' as $pb;

import 'mastermind.pbenum.dart';

export 'mastermind.pbenum.dart';

class UpdateStateRequest extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'UpdateStateRequest', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'mastermind.v1'), createEmptyInstance: create)
    ..e<StateTransition>(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'stateTransition', $pb.PbFieldType.OE, defaultOrMaker: StateTransition.STATE_TRANSITION_UNSPECIFIED, valueOf: StateTransition.valueOf, enumValues: StateTransition.values)
    ..hasRequiredFields = false
  ;

  UpdateStateRequest._() : super();
  factory UpdateStateRequest({
    StateTransition? stateTransition,
  }) {
    final _result = create();
    if (stateTransition != null) {
      _result.stateTransition = stateTransition;
    }
    return _result;
  }
  factory UpdateStateRequest.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory UpdateStateRequest.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  UpdateStateRequest clone() => UpdateStateRequest()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  UpdateStateRequest copyWith(void Function(UpdateStateRequest) updates) => super.copyWith((message) => updates(message as UpdateStateRequest)) as UpdateStateRequest; // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static UpdateStateRequest create() => UpdateStateRequest._();
  UpdateStateRequest createEmptyInstance() => create();
  static $pb.PbList<UpdateStateRequest> createRepeated() => $pb.PbList<UpdateStateRequest>();
  @$core.pragma('dart2js:noInline')
  static UpdateStateRequest getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<UpdateStateRequest>(create);
  static UpdateStateRequest? _defaultInstance;

  @$pb.TagNumber(1)
  StateTransition get stateTransition => $_getN(0);
  @$pb.TagNumber(1)
  set stateTransition(StateTransition v) { setField(1, v); }
  @$pb.TagNumber(1)
  $core.bool hasStateTransition() => $_has(0);
  @$pb.TagNumber(1)
  void clearStateTransition() => clearField(1);
}

class UpdateStateResponse extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'UpdateStateResponse', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'mastermind.v1'), createEmptyInstance: create)
    ..e<MissionState>(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'oldState', $pb.PbFieldType.OE, defaultOrMaker: MissionState.MISSION_STATE_UNSPECIFIED, valueOf: MissionState.valueOf, enumValues: MissionState.values)
    ..e<StateTransition>(2, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'stateTransition', $pb.PbFieldType.OE, defaultOrMaker: StateTransition.STATE_TRANSITION_UNSPECIFIED, valueOf: StateTransition.valueOf, enumValues: StateTransition.values)
    ..e<MissionState>(3, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'currentState', $pb.PbFieldType.OE, defaultOrMaker: MissionState.MISSION_STATE_UNSPECIFIED, valueOf: MissionState.valueOf, enumValues: MissionState.values)
    ..hasRequiredFields = false
  ;

  UpdateStateResponse._() : super();
  factory UpdateStateResponse({
    MissionState? oldState,
    StateTransition? stateTransition,
    MissionState? currentState,
  }) {
    final _result = create();
    if (oldState != null) {
      _result.oldState = oldState;
    }
    if (stateTransition != null) {
      _result.stateTransition = stateTransition;
    }
    if (currentState != null) {
      _result.currentState = currentState;
    }
    return _result;
  }
  factory UpdateStateResponse.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory UpdateStateResponse.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  UpdateStateResponse clone() => UpdateStateResponse()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  UpdateStateResponse copyWith(void Function(UpdateStateResponse) updates) => super.copyWith((message) => updates(message as UpdateStateResponse)) as UpdateStateResponse; // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static UpdateStateResponse create() => UpdateStateResponse._();
  UpdateStateResponse createEmptyInstance() => create();
  static $pb.PbList<UpdateStateResponse> createRepeated() => $pb.PbList<UpdateStateResponse>();
  @$core.pragma('dart2js:noInline')
  static UpdateStateResponse getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<UpdateStateResponse>(create);
  static UpdateStateResponse? _defaultInstance;

  @$pb.TagNumber(1)
  MissionState get oldState => $_getN(0);
  @$pb.TagNumber(1)
  set oldState(MissionState v) { setField(1, v); }
  @$pb.TagNumber(1)
  $core.bool hasOldState() => $_has(0);
  @$pb.TagNumber(1)
  void clearOldState() => clearField(1);

  @$pb.TagNumber(2)
  StateTransition get stateTransition => $_getN(1);
  @$pb.TagNumber(2)
  set stateTransition(StateTransition v) { setField(2, v); }
  @$pb.TagNumber(2)
  $core.bool hasStateTransition() => $_has(1);
  @$pb.TagNumber(2)
  void clearStateTransition() => clearField(2);

  @$pb.TagNumber(3)
  MissionState get currentState => $_getN(2);
  @$pb.TagNumber(3)
  set currentState(MissionState v) { setField(3, v); }
  @$pb.TagNumber(3)
  $core.bool hasCurrentState() => $_has(2);
  @$pb.TagNumber(3)
  void clearCurrentState() => clearField(3);
}

class Plane extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'Plane', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'mastermind.v1'), createEmptyInstance: create)
    ..a<$core.int>(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'planeId', $pb.PbFieldType.O3)
    ..a<$core.double>(2, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'lat', $pb.PbFieldType.OD)
    ..a<$core.double>(3, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'lon', $pb.PbFieldType.OD)
    ..a<$core.double>(4, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'alt', $pb.PbFieldType.OD)
    ..a<$core.double>(5, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'heading', $pb.PbFieldType.OD)
    ..a<$core.double>(6, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'speed', $pb.PbFieldType.OD)
    ..e<AgilityLevel>(7, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'agilityLevel', $pb.PbFieldType.OE, defaultOrMaker: AgilityLevel.AGILITY_LEVEL_UNSPECIFIED, valueOf: AgilityLevel.valueOf, enumValues: AgilityLevel.values)
    ..a<$core.double>(8, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'score', $pb.PbFieldType.OD)
    ..hasRequiredFields = false
  ;

  Plane._() : super();
  factory Plane({
    $core.int? planeId,
    $core.double? lat,
    $core.double? lon,
    $core.double? alt,
    $core.double? heading,
    $core.double? speed,
    AgilityLevel? agilityLevel,
    $core.double? score,
  }) {
    final _result = create();
    if (planeId != null) {
      _result.planeId = planeId;
    }
    if (lat != null) {
      _result.lat = lat;
    }
    if (lon != null) {
      _result.lon = lon;
    }
    if (alt != null) {
      _result.alt = alt;
    }
    if (heading != null) {
      _result.heading = heading;
    }
    if (speed != null) {
      _result.speed = speed;
    }
    if (agilityLevel != null) {
      _result.agilityLevel = agilityLevel;
    }
    if (score != null) {
      _result.score = score;
    }
    return _result;
  }
  factory Plane.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory Plane.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  Plane clone() => Plane()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  Plane copyWith(void Function(Plane) updates) => super.copyWith((message) => updates(message as Plane)) as Plane; // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static Plane create() => Plane._();
  Plane createEmptyInstance() => create();
  static $pb.PbList<Plane> createRepeated() => $pb.PbList<Plane>();
  @$core.pragma('dart2js:noInline')
  static Plane getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<Plane>(create);
  static Plane? _defaultInstance;

  @$pb.TagNumber(1)
  $core.int get planeId => $_getIZ(0);
  @$pb.TagNumber(1)
  set planeId($core.int v) { $_setSignedInt32(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasPlaneId() => $_has(0);
  @$pb.TagNumber(1)
  void clearPlaneId() => clearField(1);

  @$pb.TagNumber(2)
  $core.double get lat => $_getN(1);
  @$pb.TagNumber(2)
  set lat($core.double v) { $_setDouble(1, v); }
  @$pb.TagNumber(2)
  $core.bool hasLat() => $_has(1);
  @$pb.TagNumber(2)
  void clearLat() => clearField(2);

  @$pb.TagNumber(3)
  $core.double get lon => $_getN(2);
  @$pb.TagNumber(3)
  set lon($core.double v) { $_setDouble(2, v); }
  @$pb.TagNumber(3)
  $core.bool hasLon() => $_has(2);
  @$pb.TagNumber(3)
  void clearLon() => clearField(3);

  @$pb.TagNumber(4)
  $core.double get alt => $_getN(3);
  @$pb.TagNumber(4)
  set alt($core.double v) { $_setDouble(3, v); }
  @$pb.TagNumber(4)
  $core.bool hasAlt() => $_has(3);
  @$pb.TagNumber(4)
  void clearAlt() => clearField(4);

  @$pb.TagNumber(5)
  $core.double get heading => $_getN(4);
  @$pb.TagNumber(5)
  set heading($core.double v) { $_setDouble(4, v); }
  @$pb.TagNumber(5)
  $core.bool hasHeading() => $_has(4);
  @$pb.TagNumber(5)
  void clearHeading() => clearField(5);

  @$pb.TagNumber(6)
  $core.double get speed => $_getN(5);
  @$pb.TagNumber(6)
  set speed($core.double v) { $_setDouble(5, v); }
  @$pb.TagNumber(6)
  $core.bool hasSpeed() => $_has(5);
  @$pb.TagNumber(6)
  void clearSpeed() => clearField(6);

  @$pb.TagNumber(7)
  AgilityLevel get agilityLevel => $_getN(6);
  @$pb.TagNumber(7)
  set agilityLevel(AgilityLevel v) { setField(7, v); }
  @$pb.TagNumber(7)
  $core.bool hasAgilityLevel() => $_has(6);
  @$pb.TagNumber(7)
  void clearAgilityLevel() => clearField(7);

  @$pb.TagNumber(8)
  $core.double get score => $_getN(7);
  @$pb.TagNumber(8)
  set score($core.double v) { $_setDouble(7, v); }
  @$pb.TagNumber(8)
  $core.bool hasScore() => $_has(7);
  @$pb.TagNumber(8)
  void clearScore() => clearField(8);
}

class GetPlanesRequest extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'GetPlanesRequest', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'mastermind.v1'), createEmptyInstance: create)
    ..hasRequiredFields = false
  ;

  GetPlanesRequest._() : super();
  factory GetPlanesRequest() => create();
  factory GetPlanesRequest.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory GetPlanesRequest.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  GetPlanesRequest clone() => GetPlanesRequest()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  GetPlanesRequest copyWith(void Function(GetPlanesRequest) updates) => super.copyWith((message) => updates(message as GetPlanesRequest)) as GetPlanesRequest; // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static GetPlanesRequest create() => GetPlanesRequest._();
  GetPlanesRequest createEmptyInstance() => create();
  static $pb.PbList<GetPlanesRequest> createRepeated() => $pb.PbList<GetPlanesRequest>();
  @$core.pragma('dart2js:noInline')
  static GetPlanesRequest getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<GetPlanesRequest>(create);
  static GetPlanesRequest? _defaultInstance;
}

class GetPlanesResponse extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'GetPlanesResponse', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'mastermind.v1'), createEmptyInstance: create)
    ..a<$core.int>(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'selectedPlaneId', $pb.PbFieldType.O3)
    ..pc<Plane>(2, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'planes', $pb.PbFieldType.PM, subBuilder: Plane.create)
    ..hasRequiredFields = false
  ;

  GetPlanesResponse._() : super();
  factory GetPlanesResponse({
    $core.int? selectedPlaneId,
    $core.Iterable<Plane>? planes,
  }) {
    final _result = create();
    if (selectedPlaneId != null) {
      _result.selectedPlaneId = selectedPlaneId;
    }
    if (planes != null) {
      _result.planes.addAll(planes);
    }
    return _result;
  }
  factory GetPlanesResponse.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory GetPlanesResponse.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  GetPlanesResponse clone() => GetPlanesResponse()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  GetPlanesResponse copyWith(void Function(GetPlanesResponse) updates) => super.copyWith((message) => updates(message as GetPlanesResponse)) as GetPlanesResponse; // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static GetPlanesResponse create() => GetPlanesResponse._();
  GetPlanesResponse createEmptyInstance() => create();
  static $pb.PbList<GetPlanesResponse> createRepeated() => $pb.PbList<GetPlanesResponse>();
  @$core.pragma('dart2js:noInline')
  static GetPlanesResponse getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<GetPlanesResponse>(create);
  static GetPlanesResponse? _defaultInstance;

  @$pb.TagNumber(1)
  $core.int get selectedPlaneId => $_getIZ(0);
  @$pb.TagNumber(1)
  set selectedPlaneId($core.int v) { $_setSignedInt32(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasSelectedPlaneId() => $_has(0);
  @$pb.TagNumber(1)
  void clearSelectedPlaneId() => clearField(1);

  @$pb.TagNumber(2)
  $core.List<Plane> get planes => $_getList(1);
}

