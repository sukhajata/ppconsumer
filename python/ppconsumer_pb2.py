# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: ppconsumer.proto

import sys
_b=sys.version_info[0]<3 and (lambda x:x) or (lambda x:x.encode('latin1'))
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()




DESCRIPTOR = _descriptor.FileDescriptor(
  name='ppconsumer.proto',
  package='ppconsumer',
  syntax='proto3',
  serialized_options=None,
  serialized_pb=_b('\n\x10ppconsumer.proto\x12\nppconsumer\"j\n\x08\x43onsumer\x12\x10\n\x08username\x18\x01 \x01(\t\x12\r\n\x05\x65mail\x18\x02 \x01(\t\x12/\n\rinstallations\x18\x03 \x03(\x0b\x32\x18.ppconsumer.Installation\x12\x0c\n\x04type\x18\x04 \x01(\t\"3\n\x0cInstallation\x12\x10\n\x08IDNumber\x18\x01 \x01(\t\x12\x11\n\tdeviceEUI\x18\x02 \x03(\t\"#\n\x0f\x43onsumerRequest\x12\x10\n\x08username\x18\x01 \x01(\t\"*\n\x16\x43reateConsumerResponse\x12\x10\n\x08password\x18\x01 \x01(\t\"\x19\n\x08Response\x12\r\n\x05reply\x18\x01 \x01(\t2\xaa\x02\n\x0f\x43onsumerService\x12L\n\x0e\x43reateConsumer\x12\x14.ppconsumer.Consumer\x1a\".ppconsumer.CreateConsumerResponse\"\x00\x12>\n\x0eUpdateConsumer\x12\x14.ppconsumer.Consumer\x1a\x14.ppconsumer.Consumer\"\x00\x12\x42\n\x0bGetConsumer\x12\x1b.ppconsumer.ConsumerRequest\x1a\x14.ppconsumer.Consumer\"\x00\x12\x45\n\x0e\x44\x65leteConsumer\x12\x1b.ppconsumer.ConsumerRequest\x1a\x14.ppconsumer.Response\"\x00\x62\x06proto3')
)




_CONSUMER = _descriptor.Descriptor(
  name='Consumer',
  full_name='ppconsumer.Consumer',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='username', full_name='ppconsumer.Consumer.username', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='email', full_name='ppconsumer.Consumer.email', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='installations', full_name='ppconsumer.Consumer.installations', index=2,
      number=3, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='type', full_name='ppconsumer.Consumer.type', index=3,
      number=4, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=32,
  serialized_end=138,
)


_INSTALLATION = _descriptor.Descriptor(
  name='Installation',
  full_name='ppconsumer.Installation',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='IDNumber', full_name='ppconsumer.Installation.IDNumber', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='deviceEUI', full_name='ppconsumer.Installation.deviceEUI', index=1,
      number=2, type=9, cpp_type=9, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=140,
  serialized_end=191,
)


_CONSUMERREQUEST = _descriptor.Descriptor(
  name='ConsumerRequest',
  full_name='ppconsumer.ConsumerRequest',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='username', full_name='ppconsumer.ConsumerRequest.username', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=193,
  serialized_end=228,
)


_CREATECONSUMERRESPONSE = _descriptor.Descriptor(
  name='CreateConsumerResponse',
  full_name='ppconsumer.CreateConsumerResponse',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='password', full_name='ppconsumer.CreateConsumerResponse.password', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=230,
  serialized_end=272,
)


_RESPONSE = _descriptor.Descriptor(
  name='Response',
  full_name='ppconsumer.Response',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='reply', full_name='ppconsumer.Response.reply', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=274,
  serialized_end=299,
)

_CONSUMER.fields_by_name['installations'].message_type = _INSTALLATION
DESCRIPTOR.message_types_by_name['Consumer'] = _CONSUMER
DESCRIPTOR.message_types_by_name['Installation'] = _INSTALLATION
DESCRIPTOR.message_types_by_name['ConsumerRequest'] = _CONSUMERREQUEST
DESCRIPTOR.message_types_by_name['CreateConsumerResponse'] = _CREATECONSUMERRESPONSE
DESCRIPTOR.message_types_by_name['Response'] = _RESPONSE
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

Consumer = _reflection.GeneratedProtocolMessageType('Consumer', (_message.Message,), {
  'DESCRIPTOR' : _CONSUMER,
  '__module__' : 'ppconsumer_pb2'
  # @@protoc_insertion_point(class_scope:ppconsumer.Consumer)
  })
_sym_db.RegisterMessage(Consumer)

Installation = _reflection.GeneratedProtocolMessageType('Installation', (_message.Message,), {
  'DESCRIPTOR' : _INSTALLATION,
  '__module__' : 'ppconsumer_pb2'
  # @@protoc_insertion_point(class_scope:ppconsumer.Installation)
  })
_sym_db.RegisterMessage(Installation)

ConsumerRequest = _reflection.GeneratedProtocolMessageType('ConsumerRequest', (_message.Message,), {
  'DESCRIPTOR' : _CONSUMERREQUEST,
  '__module__' : 'ppconsumer_pb2'
  # @@protoc_insertion_point(class_scope:ppconsumer.ConsumerRequest)
  })
_sym_db.RegisterMessage(ConsumerRequest)

CreateConsumerResponse = _reflection.GeneratedProtocolMessageType('CreateConsumerResponse', (_message.Message,), {
  'DESCRIPTOR' : _CREATECONSUMERRESPONSE,
  '__module__' : 'ppconsumer_pb2'
  # @@protoc_insertion_point(class_scope:ppconsumer.CreateConsumerResponse)
  })
_sym_db.RegisterMessage(CreateConsumerResponse)

Response = _reflection.GeneratedProtocolMessageType('Response', (_message.Message,), {
  'DESCRIPTOR' : _RESPONSE,
  '__module__' : 'ppconsumer_pb2'
  # @@protoc_insertion_point(class_scope:ppconsumer.Response)
  })
_sym_db.RegisterMessage(Response)



_CONSUMERSERVICE = _descriptor.ServiceDescriptor(
  name='ConsumerService',
  full_name='ppconsumer.ConsumerService',
  file=DESCRIPTOR,
  index=0,
  serialized_options=None,
  serialized_start=302,
  serialized_end=600,
  methods=[
  _descriptor.MethodDescriptor(
    name='CreateConsumer',
    full_name='ppconsumer.ConsumerService.CreateConsumer',
    index=0,
    containing_service=None,
    input_type=_CONSUMER,
    output_type=_CREATECONSUMERRESPONSE,
    serialized_options=None,
  ),
  _descriptor.MethodDescriptor(
    name='UpdateConsumer',
    full_name='ppconsumer.ConsumerService.UpdateConsumer',
    index=1,
    containing_service=None,
    input_type=_CONSUMER,
    output_type=_CONSUMER,
    serialized_options=None,
  ),
  _descriptor.MethodDescriptor(
    name='GetConsumer',
    full_name='ppconsumer.ConsumerService.GetConsumer',
    index=2,
    containing_service=None,
    input_type=_CONSUMERREQUEST,
    output_type=_CONSUMER,
    serialized_options=None,
  ),
  _descriptor.MethodDescriptor(
    name='DeleteConsumer',
    full_name='ppconsumer.ConsumerService.DeleteConsumer',
    index=3,
    containing_service=None,
    input_type=_CONSUMERREQUEST,
    output_type=_RESPONSE,
    serialized_options=None,
  ),
])
_sym_db.RegisterServiceDescriptor(_CONSUMERSERVICE)

DESCRIPTOR.services_by_name['ConsumerService'] = _CONSUMERSERVICE

# @@protoc_insertion_point(module_scope)
