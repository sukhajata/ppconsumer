/**
 * @fileoverview gRPC-Web generated client stub for pplogger
 * @enhanceable
 * @public
 */

// GENERATED CODE -- DO NOT EDIT!



const grpc = {};
grpc.web = require('grpc-web');

const proto = {};
proto.pplogger = require('./ppconsumer_pb.js');

/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?Object} options
 * @constructor
 * @struct
 * @final
 */
proto.pplogger.LoggerServiceClient =
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
proto.pplogger.LoggerServicePromiseClient =
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
 *   !proto.pplogger.Consumer,
 *   !proto.pplogger.Consumer>}
 */
const methodDescriptor_LoggerService_CreateConsumer = new grpc.web.MethodDescriptor(
  '/pplogger.LoggerService/CreateConsumer',
  grpc.web.MethodType.UNARY,
  proto.pplogger.Consumer,
  proto.pplogger.Consumer,
  /**
   * @param {!proto.pplogger.Consumer} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.pplogger.Consumer.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.pplogger.Consumer,
 *   !proto.pplogger.Consumer>}
 */
const methodInfo_LoggerService_CreateConsumer = new grpc.web.AbstractClientBase.MethodInfo(
  proto.pplogger.Consumer,
  /**
   * @param {!proto.pplogger.Consumer} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.pplogger.Consumer.deserializeBinary
);


/**
 * @param {!proto.pplogger.Consumer} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.pplogger.Consumer)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.pplogger.Consumer>|undefined}
 *     The XHR Node Readable Stream
 */
proto.pplogger.LoggerServiceClient.prototype.createConsumer =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/pplogger.LoggerService/CreateConsumer',
      request,
      metadata || {},
      methodDescriptor_LoggerService_CreateConsumer,
      callback);
};


/**
 * @param {!proto.pplogger.Consumer} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.pplogger.Consumer>}
 *     A native promise that resolves to the response
 */
proto.pplogger.LoggerServicePromiseClient.prototype.createConsumer =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/pplogger.LoggerService/CreateConsumer',
      request,
      metadata || {},
      methodDescriptor_LoggerService_CreateConsumer);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.pplogger.Consumer,
 *   !proto.pplogger.Consumer>}
 */
const methodDescriptor_LoggerService_UpdateConsumer = new grpc.web.MethodDescriptor(
  '/pplogger.LoggerService/UpdateConsumer',
  grpc.web.MethodType.UNARY,
  proto.pplogger.Consumer,
  proto.pplogger.Consumer,
  /**
   * @param {!proto.pplogger.Consumer} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.pplogger.Consumer.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.pplogger.Consumer,
 *   !proto.pplogger.Consumer>}
 */
const methodInfo_LoggerService_UpdateConsumer = new grpc.web.AbstractClientBase.MethodInfo(
  proto.pplogger.Consumer,
  /**
   * @param {!proto.pplogger.Consumer} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.pplogger.Consumer.deserializeBinary
);


/**
 * @param {!proto.pplogger.Consumer} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.pplogger.Consumer)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.pplogger.Consumer>|undefined}
 *     The XHR Node Readable Stream
 */
proto.pplogger.LoggerServiceClient.prototype.updateConsumer =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/pplogger.LoggerService/UpdateConsumer',
      request,
      metadata || {},
      methodDescriptor_LoggerService_UpdateConsumer,
      callback);
};


/**
 * @param {!proto.pplogger.Consumer} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.pplogger.Consumer>}
 *     A native promise that resolves to the response
 */
proto.pplogger.LoggerServicePromiseClient.prototype.updateConsumer =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/pplogger.LoggerService/UpdateConsumer',
      request,
      metadata || {},
      methodDescriptor_LoggerService_UpdateConsumer);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.pplogger.ConsumerRequest,
 *   !proto.pplogger.Consumer>}
 */
const methodDescriptor_LoggerService_GetConsumer = new grpc.web.MethodDescriptor(
  '/pplogger.LoggerService/GetConsumer',
  grpc.web.MethodType.UNARY,
  proto.pplogger.ConsumerRequest,
  proto.pplogger.Consumer,
  /**
   * @param {!proto.pplogger.ConsumerRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.pplogger.Consumer.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.pplogger.ConsumerRequest,
 *   !proto.pplogger.Consumer>}
 */
const methodInfo_LoggerService_GetConsumer = new grpc.web.AbstractClientBase.MethodInfo(
  proto.pplogger.Consumer,
  /**
   * @param {!proto.pplogger.ConsumerRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.pplogger.Consumer.deserializeBinary
);


/**
 * @param {!proto.pplogger.ConsumerRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.pplogger.Consumer)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.pplogger.Consumer>|undefined}
 *     The XHR Node Readable Stream
 */
proto.pplogger.LoggerServiceClient.prototype.getConsumer =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/pplogger.LoggerService/GetConsumer',
      request,
      metadata || {},
      methodDescriptor_LoggerService_GetConsumer,
      callback);
};


/**
 * @param {!proto.pplogger.ConsumerRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.pplogger.Consumer>}
 *     A native promise that resolves to the response
 */
proto.pplogger.LoggerServicePromiseClient.prototype.getConsumer =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/pplogger.LoggerService/GetConsumer',
      request,
      metadata || {},
      methodDescriptor_LoggerService_GetConsumer);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.pplogger.ConsumerRequest,
 *   !proto.pplogger.Response>}
 */
const methodDescriptor_LoggerService_DeleteConsumer = new grpc.web.MethodDescriptor(
  '/pplogger.LoggerService/DeleteConsumer',
  grpc.web.MethodType.UNARY,
  proto.pplogger.ConsumerRequest,
  proto.pplogger.Response,
  /**
   * @param {!proto.pplogger.ConsumerRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.pplogger.Response.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.pplogger.ConsumerRequest,
 *   !proto.pplogger.Response>}
 */
const methodInfo_LoggerService_DeleteConsumer = new grpc.web.AbstractClientBase.MethodInfo(
  proto.pplogger.Response,
  /**
   * @param {!proto.pplogger.ConsumerRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.pplogger.Response.deserializeBinary
);


/**
 * @param {!proto.pplogger.ConsumerRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.pplogger.Response)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.pplogger.Response>|undefined}
 *     The XHR Node Readable Stream
 */
proto.pplogger.LoggerServiceClient.prototype.deleteConsumer =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/pplogger.LoggerService/DeleteConsumer',
      request,
      metadata || {},
      methodDescriptor_LoggerService_DeleteConsumer,
      callback);
};


/**
 * @param {!proto.pplogger.ConsumerRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.pplogger.Response>}
 *     A native promise that resolves to the response
 */
proto.pplogger.LoggerServicePromiseClient.prototype.deleteConsumer =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/pplogger.LoggerService/DeleteConsumer',
      request,
      metadata || {},
      methodDescriptor_LoggerService_DeleteConsumer);
};


module.exports = proto.pplogger;

