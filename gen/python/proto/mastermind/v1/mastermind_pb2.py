# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: proto/mastermind/v1/mastermind.proto
"""Generated protocol buffer code."""
from google.protobuf.internal import builder as _builder
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()




DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n$proto/mastermind/v1/mastermind.proto\x12\rmastermind.v1\"\xa7\x01\n\x12UpdateStateRequest\x12\x46\n\x0fsending_service\x18\x01 \x01(\x0e\x32\x1d.mastermind.v1.SendingServiceR\x0esendingService\x12I\n\x10state_transition\x18\x02 \x01(\x0e\x32\x1e.mastermind.v1.StateTransitionR\x0fstateTransition\"\xdc\x01\n\x13UpdateStateResponse\x12\x38\n\told_state\x18\x01 \x01(\x0e\x32\x1b.mastermind.v1.MissionStateR\x08oldState\x12I\n\x10state_transition\x18\x02 \x01(\x0e\x32\x1e.mastermind.v1.StateTransitionR\x0fstateTransition\x12@\n\rcurrent_state\x18\x03 \x01(\x0e\x32\x1b.mastermind.v1.MissionStateR\x0c\x63urrentState*\x82\x01\n\x0cMissionState\x12\x1d\n\x19MISSION_STATE_UNSPECIFIED\x10\x00\x12\x1a\n\x16MISSION_STATE_APPROACH\x10\x01\x12\x1b\n\x17MISSION_STATE_FOLLOWING\x10\x02\x12\x1a\n\x16MISSION_STATE_KAMIKAZE\x10\x03*\xbb\x02\n\x0fStateTransition\x12 \n\x1cSTATE_TRANSITION_UNSPECIFIED\x10\x00\x12&\n\"STATE_TRANSITION_TARGET_APPROACHED\x10\x01\x12 \n\x1cSTATE_TRANSITION_LOCK_FAILED\x10\x02\x12!\n\x1dSTATE_TRANSITION_LOCK_SUCCESS\x10\x03\x12\x1f\n\x1bSTATE_TRANSITION_QR_SUCCESS\x10\x04\x12\x1e\n\x1aSTATE_TRANSITION_QR_FAILED\x10\x05\x12+\n\'STATE_TRANSITION_MODE_KAMIKAZE_SELECTED\x10\x06\x12+\n\'STATE_TRANSITION_MODE_APPROACH_SELECTED\x10\x07*\xa0\x01\n\x0eSendingService\x12\x1f\n\x1bSENDING_SERVICE_UNSPECIFIED\x10\x00\x12\x1c\n\x18SENDING_SERVICE_APPROACH\x10\x01\x12\x18\n\x14SENDING_SERVICE_LOCK\x10\x02\x12\x1c\n\x18SENDING_SERVICE_KAMIKAZE\x10\x03\x12\x17\n\x13SENDING_SERVICE_GUI\x10\x04\x32m\n\x11MastermindService\x12X\n\x0bUpdateState\x12!.mastermind.v1.UpdateStateRequest\x1a\".mastermind.v1.UpdateStateResponse(\x01\x30\x01\x42\xc7\x01\n\x11\x63om.mastermind.v1B\x0fMastermindProtoP\x01ZLgithub.com/koustech/mastermind/gen/proto/go/proto/mastermind/v1;mastermindv1\xa2\x02\x03MXX\xaa\x02\rMastermind.V1\xca\x02\rMastermind\\V1\xe2\x02\x19Mastermind\\V1\\GPBMetadata\xea\x02\x0eMastermind::V1b\x06proto3')

_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, globals())
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'proto.mastermind.v1.mastermind_pb2', globals())
if _descriptor._USE_C_DESCRIPTORS == False:

  DESCRIPTOR._options = None
  DESCRIPTOR._serialized_options = b'\n\021com.mastermind.v1B\017MastermindProtoP\001ZLgithub.com/koustech/mastermind/gen/proto/go/proto/mastermind/v1;mastermindv1\242\002\003MXX\252\002\rMastermind.V1\312\002\rMastermind\\V1\342\002\031Mastermind\\V1\\GPBMetadata\352\002\016Mastermind::V1'
  _MISSIONSTATE._serialized_start=449
  _MISSIONSTATE._serialized_end=579
  _STATETRANSITION._serialized_start=582
  _STATETRANSITION._serialized_end=897
  _SENDINGSERVICE._serialized_start=900
  _SENDINGSERVICE._serialized_end=1060
  _UPDATESTATEREQUEST._serialized_start=56
  _UPDATESTATEREQUEST._serialized_end=223
  _UPDATESTATERESPONSE._serialized_start=226
  _UPDATESTATERESPONSE._serialized_end=446
  _MASTERMINDSERVICE._serialized_start=1062
  _MASTERMINDSERVICE._serialized_end=1171
# @@protoc_insertion_point(module_scope)
