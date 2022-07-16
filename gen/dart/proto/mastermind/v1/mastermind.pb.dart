///
//  Generated code. Do not modify.
//  source: proto/mastermind/v1/mastermind.proto
//
// @dart = 2.12
// ignore_for_file: annotate_overrides,camel_case_types,constant_identifier_names,directives_ordering,library_prefixes,non_constant_identifier_names,prefer_final_fields,return_of_invalid_type,unnecessary_const,unnecessary_import,unnecessary_this,unused_import,unused_shown_name

import 'dart:async' as $async;
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

class MastermindServiceApi {
  $pb.RpcClient _client;
  MastermindServiceApi(this._client);

  $async.Future<UpdateStateResponse> updateState($pb.ClientContext? ctx, UpdateStateRequest request) {
    var emptyResponse = UpdateStateResponse();
    return _client.invoke<UpdateStateResponse>(ctx, 'MastermindService', 'UpdateState', request, emptyResponse);
  }
}

