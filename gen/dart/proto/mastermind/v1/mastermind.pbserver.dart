///
//  Generated code. Do not modify.
//  source: proto/mastermind/v1/mastermind.proto
//
// @dart = 2.12
// ignore_for_file: annotate_overrides,camel_case_types,constant_identifier_names,deprecated_member_use_from_same_package,directives_ordering,library_prefixes,non_constant_identifier_names,prefer_final_fields,return_of_invalid_type,unnecessary_const,unnecessary_import,unnecessary_this,unused_import,unused_shown_name

import 'dart:async' as $async;

import 'package:protobuf/protobuf.dart' as $pb;

import 'dart:core' as $core;
import 'mastermind.pb.dart' as $0;
import 'mastermind.pbjson.dart';

export 'mastermind.pb.dart';

abstract class MastermindServiceBase extends $pb.GeneratedService {
  $async.Future<$0.UpdateStateResponse> updateState($pb.ServerContext ctx, $0.UpdateStateRequest request);

  $pb.GeneratedMessage createRequest($core.String method) {
    switch (method) {
      case 'UpdateState': return $0.UpdateStateRequest();
      default: throw $core.ArgumentError('Unknown method: $method');
    }
  }

  $async.Future<$pb.GeneratedMessage> handleCall($pb.ServerContext ctx, $core.String method, $pb.GeneratedMessage request) {
    switch (method) {
      case 'UpdateState': return this.updateState(ctx, request as $0.UpdateStateRequest);
      default: throw $core.ArgumentError('Unknown method: $method');
    }
  }

  $core.Map<$core.String, $core.dynamic> get $json => MastermindServiceBase$json;
  $core.Map<$core.String, $core.Map<$core.String, $core.dynamic>> get $messageJson => MastermindServiceBase$messageJson;
}

