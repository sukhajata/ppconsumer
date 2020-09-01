/**
 * @fileoverview gRPC-Web generated client stub for ppconsumer
 * @enhanceable
 * @public
 */

// GENERATED CODE -- DO NOT EDIT!



const grpc = {};
grpc.web = require('grpc-web');

const proto = {};
proto.ppconsumer = require('./ppconsumer_pb.js');

/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?Object} options
 * @constructor
 * @struct
 * @final
 */
proto.ppconsumer.ConsumerServiceClient =
    function(hostname, credentials, options) {
  if (!options) options = {};
  options['format'] = 'text';

  /**
   * @private @const {!grpc.web.GrpcWebClientBase} The client
   */
  this.client_ = new grpc.web.GrpcWebClientBase(options);

  /**
   * @private @const {string} The hostname
   */
  this.hostname_ = hostname;

};


/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?Object} options
 * @constructor
 * @struct
 * @final
 */
proto.ppconsumer.ConsumerServicePromiseClient =
    function(hostname, credentials, options) {
  if (!options) options = {};
  options['format'] = 'text';

  /**
   * @private @const {!grpc.web.GrpcWebClientBase} The client
   */
  this.client_ = new grpc.web.GrpcWebClientBase(options);

  /**
   * @private @const {string} The hostname
   */
  this.hostname_ = hostname;

};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.ppconsumer.Consumer,
 *   !proto.ppconsumer.Consumer>}
 */
const methodDescriptor_ConsumerService_CreateConsumer = new grpc.web.MethodDescriptor(
  '/ppconsumer.ConsumerService/CreateConsumer',
  grpc.web.MethodType.UNARY,
  proto.ppconsumer.Consumer,
  proto.ppconsumer.Consumer,
  /**
   * @param {!proto.ppconsumer.Consumer} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.ppconsumer.Consumer.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.ppconsumer.Consumer,
 *   !proto.ppconsumer.Consumer>}
 */
const methodInfo_ConsumerService_CreateConsumer = new grpc.web.AbstractClientBase.MethodInfo(
  proto.ppconsumer.Consumer,
  /**
   * @param {!proto.ppconsumer.Consumer} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.ppconsumer.Consumer.deserializeBinary
);


/**
 * @param {!proto.ppconsumer.Consumer} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.ppconsumer.Consumer)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.ppconsumer.Consumer>|undefined}
 *     The XHR Node Readable Stream
 */
proto.ppconsumer.ConsumerServiceClient.prototype.createConsumer =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/ppconsumer.ConsumerService/CreateConsumer',
      request,
      metadata || {},
      methodDescriptor_ConsumerService_CreateConsumer,
      callback);
};


/**
 * @param {!proto.ppconsumer.Consumer} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.ppconsumer.Consumer>}
 *     A native promise that resolves to the response
 */
proto.ppconsumer.ConsumerServicePromiseClient.prototype.createConsumer =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/ppconsumer.ConsumerService/CreateConsumer',
      request,
      metadata || {},
      methodDescriptor_ConsumerService_CreateConsumer);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.ppconsumer.Consumer,
 *   !proto.ppconsumer.Consumer>}
 */
const methodDescriptor_ConsumerService_UpdateConsumer = new grpc.web.MethodDescriptor(
  '/ppconsumer.ConsumerService/UpdateConsumer',
  grpc.web.MethodType.UNARY,
  proto.ppconsumer.Consumer,
  proto.ppconsumer.Consumer,
  /**
   * @param {!proto.ppconsumer.Consumer} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.ppconsumer.Consumer.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.ppconsumer.Consumer,
 *   !proto.ppconsumer.Consumer>}
 */
const methodInfo_ConsumerService_UpdateConsumer = new grpc.web.AbstractClientBase.MethodInfo(
  proto.ppconsumer.Consumer,
  /**
   * @param {!proto.ppconsumer.Consumer} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.ppconsumer.Consumer.deserializeBinary
);


/**
 * @param {!proto.ppconsumer.Consumer} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.ppconsumer.Consumer)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.ppconsumer.Consumer>|undefined}
 *     The XHR Node Readable Stream
 */
proto.ppconsumer.ConsumerServiceClient.prototype.updateConsumer =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/ppconsumer.ConsumerService/UpdateConsumer',
      request,
      metadata || {},
      methodDescriptor_ConsumerService_UpdateConsumer,
      callback);
};


/**
 * @param {!proto.ppconsumer.Consumer} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.ppconsumer.Consumer>}
 *     A native promise that resolves to the response
 */
proto.ppconsumer.ConsumerServicePromiseClient.prototype.updateConsumer =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/ppconsumer.ConsumerService/UpdateConsumer',
      request,
      metadata || {},
      methodDescriptor_ConsumerService_UpdateConsumer);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.ppconsumer.ConsumerRequest,
 *   !proto.ppconsumer.Consumer>}
 */
const methodDescriptor_ConsumerService_GetConsumer = new grpc.web.MethodDescriptor(
  '/ppconsumer.ConsumerService/GetConsumer',
  grpc.web.MethodType.UNARY,
  proto.ppconsumer.ConsumerRequest,
  proto.ppconsumer.Consumer,
  /**
   * @param {!proto.ppconsumer.ConsumerRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.ppconsumer.Consumer.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.ppconsumer.ConsumerRequest,
 *   !proto.ppconsumer.Consumer>}
 */
const methodInfo_ConsumerService_GetConsumer = new grpc.web.AbstractClientBase.MethodInfo(
  proto.ppconsumer.Consumer,
  /**
   * @param {!proto.ppconsumer.ConsumerRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.ppconsumer.Consumer.deserializeBinary
);


/**
 * @param {!proto.ppconsumer.ConsumerRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.ppconsumer.Consumer)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.ppconsumer.Consumer>|undefined}
 *     The XHR Node Readable Stream
 */
proto.ppconsumer.ConsumerServiceClient.prototype.getConsumer =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/ppconsumer.ConsumerService/GetConsumer',
      request,
      metadata || {},
      methodDescriptor_ConsumerService_GetConsumer,
      callback);
};


/**
 * @param {!proto.ppconsumer.ConsumerRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.ppconsumer.Consumer>}
 *     A native promise that resolves to the response
 */
proto.ppconsumer.ConsumerServicePromiseClient.prototype.getConsumer =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/ppconsumer.ConsumerService/GetConsumer',
      request,
      metadata || {},
      methodDescriptor_ConsumerService_GetConsumer);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.ppconsumer.ConsumerRequest,
 *   !proto.ppconsumer.Response>}
 */
const methodDescriptor_ConsumerService_DeleteConsumer = new grpc.web.MethodDescriptor(
  '/ppconsumer.ConsumerService/DeleteConsumer',
  grpc.web.MethodType.UNARY,
  proto.ppconsumer.ConsumerRequest,
  proto.ppconsumer.Response,
  /**
   * @param {!proto.ppconsumer.ConsumerRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.ppconsumer.Response.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.ppconsumer.ConsumerRequest,
 *   !proto.ppconsumer.Response>}
 */
const methodInfo_ConsumerService_DeleteConsumer = new grpc.web.AbstractClientBase.MethodInfo(
  proto.ppconsumer.Response,
  /**
   * @param {!proto.ppconsumer.ConsumerRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.ppconsumer.Response.deserializeBinary
);


/**
 * @param {!proto.ppconsumer.ConsumerRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.ppconsumer.Response)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.ppconsumer.Response>|undefined}
 *     The XHR Node Readable Stream
 */
proto.ppconsumer.ConsumerServiceClient.prototype.deleteConsumer =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/ppconsumer.ConsumerService/DeleteConsumer',
      request,
      metadata || {},
      methodDescriptor_ConsumerService_DeleteConsumer,
      callback);
};


/**
 * @param {!proto.ppconsumer.ConsumerRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.ppconsumer.Response>}
 *     A native promise that resolves to the response
 */
proto.ppconsumer.ConsumerServicePromiseClient.prototype.deleteConsumer =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/ppconsumer.ConsumerService/DeleteConsumer',
      request,
      metadata || {},
      methodDescriptor_ConsumerService_DeleteConsumer);
};


module.exports = proto.ppconsumer;

