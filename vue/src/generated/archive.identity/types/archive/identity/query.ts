/* eslint-disable */
import Long from "long";
import _m0 from "protobufjs/minimal";
import { PageRequest, PageResponse } from "../../cosmos/base/query/v1beta1/pagination";
import { Certificate } from "./certificate";
import { Issuer } from "./issuer";
import { Params } from "./params";

export const protobufPackage = "archive.identity";

/** QueryParamsRequest is request type for the Query/Params RPC method. */
export interface QueryParamsRequest {
}

/** QueryParamsResponse is response type for the Query/Params RPC method. */
export interface QueryParamsResponse {
  /** params holds all the parameters of this module. */
  params: Params | undefined;
}

export interface QueryIdentityMembersRequest {
  id: number;
  isPending: boolean;
  pagination: PageRequest | undefined;
}

export interface QueryIdentityMembersResponse {
  members: string[];
  pagination: PageResponse | undefined;
}

export interface QueryIssuersRequest {
  pagination: PageRequest | undefined;
}

export interface QueryIssuersResponse {
  issuers: string[];
  pagination: PageResponse | undefined;
}

export interface QueryIssuerInfoRequest {
  issuer: string;
}

export interface QueryIssuerInfoResponse {
  issuerInfo: Issuer | undefined;
}

export interface QueryIdentityRequest {
  id: number;
}

export interface QueryIdentityResponse {
  certificate: Certificate | undefined;
}

export interface QueryOperatorsRequest {
  id: number;
  pagination: PageRequest | undefined;
}

export interface QueryOperatorsResponse {
  pagination: PageResponse | undefined;
}

function createBaseQueryParamsRequest(): QueryParamsRequest {
  return {};
}

export const QueryParamsRequest = {
  encode(_: QueryParamsRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryParamsRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryParamsRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(_: any): QueryParamsRequest {
    return {};
  },

  toJSON(_: QueryParamsRequest): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryParamsRequest>, I>>(_: I): QueryParamsRequest {
    const message = createBaseQueryParamsRequest();
    return message;
  },
};

function createBaseQueryParamsResponse(): QueryParamsResponse {
  return { params: undefined };
}

export const QueryParamsResponse = {
  encode(message: QueryParamsResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.params !== undefined) {
      Params.encode(message.params, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryParamsResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryParamsResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.params = Params.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryParamsResponse {
    return { params: isSet(object.params) ? Params.fromJSON(object.params) : undefined };
  },

  toJSON(message: QueryParamsResponse): unknown {
    const obj: any = {};
    message.params !== undefined && (obj.params = message.params ? Params.toJSON(message.params) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryParamsResponse>, I>>(object: I): QueryParamsResponse {
    const message = createBaseQueryParamsResponse();
    message.params = (object.params !== undefined && object.params !== null)
      ? Params.fromPartial(object.params)
      : undefined;
    return message;
  },
};

function createBaseQueryIdentityMembersRequest(): QueryIdentityMembersRequest {
  return { id: 0, isPending: false, pagination: undefined };
}

export const QueryIdentityMembersRequest = {
  encode(message: QueryIdentityMembersRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.id !== 0) {
      writer.uint32(8).uint64(message.id);
    }
    if (message.isPending === true) {
      writer.uint32(16).bool(message.isPending);
    }
    if (message.pagination !== undefined) {
      PageRequest.encode(message.pagination, writer.uint32(26).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryIdentityMembersRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryIdentityMembersRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.id = longToNumber(reader.uint64() as Long);
          break;
        case 2:
          message.isPending = reader.bool();
          break;
        case 3:
          message.pagination = PageRequest.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryIdentityMembersRequest {
    return {
      id: isSet(object.id) ? Number(object.id) : 0,
      isPending: isSet(object.isPending) ? Boolean(object.isPending) : false,
      pagination: isSet(object.pagination) ? PageRequest.fromJSON(object.pagination) : undefined,
    };
  },

  toJSON(message: QueryIdentityMembersRequest): unknown {
    const obj: any = {};
    message.id !== undefined && (obj.id = Math.round(message.id));
    message.isPending !== undefined && (obj.isPending = message.isPending);
    message.pagination !== undefined
      && (obj.pagination = message.pagination ? PageRequest.toJSON(message.pagination) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryIdentityMembersRequest>, I>>(object: I): QueryIdentityMembersRequest {
    const message = createBaseQueryIdentityMembersRequest();
    message.id = object.id ?? 0;
    message.isPending = object.isPending ?? false;
    message.pagination = (object.pagination !== undefined && object.pagination !== null)
      ? PageRequest.fromPartial(object.pagination)
      : undefined;
    return message;
  },
};

function createBaseQueryIdentityMembersResponse(): QueryIdentityMembersResponse {
  return { members: [], pagination: undefined };
}

export const QueryIdentityMembersResponse = {
  encode(message: QueryIdentityMembersResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    for (const v of message.members) {
      writer.uint32(10).string(v!);
    }
    if (message.pagination !== undefined) {
      PageResponse.encode(message.pagination, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryIdentityMembersResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryIdentityMembersResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.members.push(reader.string());
          break;
        case 2:
          message.pagination = PageResponse.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryIdentityMembersResponse {
    return {
      members: Array.isArray(object?.members) ? object.members.map((e: any) => String(e)) : [],
      pagination: isSet(object.pagination) ? PageResponse.fromJSON(object.pagination) : undefined,
    };
  },

  toJSON(message: QueryIdentityMembersResponse): unknown {
    const obj: any = {};
    if (message.members) {
      obj.members = message.members.map((e) => e);
    } else {
      obj.members = [];
    }
    message.pagination !== undefined
      && (obj.pagination = message.pagination ? PageResponse.toJSON(message.pagination) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryIdentityMembersResponse>, I>>(object: I): QueryIdentityMembersResponse {
    const message = createBaseQueryIdentityMembersResponse();
    message.members = object.members?.map((e) => e) || [];
    message.pagination = (object.pagination !== undefined && object.pagination !== null)
      ? PageResponse.fromPartial(object.pagination)
      : undefined;
    return message;
  },
};

function createBaseQueryIssuersRequest(): QueryIssuersRequest {
  return { pagination: undefined };
}

export const QueryIssuersRequest = {
  encode(message: QueryIssuersRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.pagination !== undefined) {
      PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryIssuersRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryIssuersRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.pagination = PageRequest.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryIssuersRequest {
    return { pagination: isSet(object.pagination) ? PageRequest.fromJSON(object.pagination) : undefined };
  },

  toJSON(message: QueryIssuersRequest): unknown {
    const obj: any = {};
    message.pagination !== undefined
      && (obj.pagination = message.pagination ? PageRequest.toJSON(message.pagination) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryIssuersRequest>, I>>(object: I): QueryIssuersRequest {
    const message = createBaseQueryIssuersRequest();
    message.pagination = (object.pagination !== undefined && object.pagination !== null)
      ? PageRequest.fromPartial(object.pagination)
      : undefined;
    return message;
  },
};

function createBaseQueryIssuersResponse(): QueryIssuersResponse {
  return { issuers: [], pagination: undefined };
}

export const QueryIssuersResponse = {
  encode(message: QueryIssuersResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    for (const v of message.issuers) {
      writer.uint32(10).string(v!);
    }
    if (message.pagination !== undefined) {
      PageResponse.encode(message.pagination, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryIssuersResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryIssuersResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.issuers.push(reader.string());
          break;
        case 2:
          message.pagination = PageResponse.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryIssuersResponse {
    return {
      issuers: Array.isArray(object?.issuers) ? object.issuers.map((e: any) => String(e)) : [],
      pagination: isSet(object.pagination) ? PageResponse.fromJSON(object.pagination) : undefined,
    };
  },

  toJSON(message: QueryIssuersResponse): unknown {
    const obj: any = {};
    if (message.issuers) {
      obj.issuers = message.issuers.map((e) => e);
    } else {
      obj.issuers = [];
    }
    message.pagination !== undefined
      && (obj.pagination = message.pagination ? PageResponse.toJSON(message.pagination) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryIssuersResponse>, I>>(object: I): QueryIssuersResponse {
    const message = createBaseQueryIssuersResponse();
    message.issuers = object.issuers?.map((e) => e) || [];
    message.pagination = (object.pagination !== undefined && object.pagination !== null)
      ? PageResponse.fromPartial(object.pagination)
      : undefined;
    return message;
  },
};

function createBaseQueryIssuerInfoRequest(): QueryIssuerInfoRequest {
  return { issuer: "" };
}

export const QueryIssuerInfoRequest = {
  encode(message: QueryIssuerInfoRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.issuer !== "") {
      writer.uint32(10).string(message.issuer);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryIssuerInfoRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryIssuerInfoRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.issuer = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryIssuerInfoRequest {
    return { issuer: isSet(object.issuer) ? String(object.issuer) : "" };
  },

  toJSON(message: QueryIssuerInfoRequest): unknown {
    const obj: any = {};
    message.issuer !== undefined && (obj.issuer = message.issuer);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryIssuerInfoRequest>, I>>(object: I): QueryIssuerInfoRequest {
    const message = createBaseQueryIssuerInfoRequest();
    message.issuer = object.issuer ?? "";
    return message;
  },
};

function createBaseQueryIssuerInfoResponse(): QueryIssuerInfoResponse {
  return { issuerInfo: undefined };
}

export const QueryIssuerInfoResponse = {
  encode(message: QueryIssuerInfoResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.issuerInfo !== undefined) {
      Issuer.encode(message.issuerInfo, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryIssuerInfoResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryIssuerInfoResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.issuerInfo = Issuer.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryIssuerInfoResponse {
    return { issuerInfo: isSet(object.issuerInfo) ? Issuer.fromJSON(object.issuerInfo) : undefined };
  },

  toJSON(message: QueryIssuerInfoResponse): unknown {
    const obj: any = {};
    message.issuerInfo !== undefined
      && (obj.issuerInfo = message.issuerInfo ? Issuer.toJSON(message.issuerInfo) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryIssuerInfoResponse>, I>>(object: I): QueryIssuerInfoResponse {
    const message = createBaseQueryIssuerInfoResponse();
    message.issuerInfo = (object.issuerInfo !== undefined && object.issuerInfo !== null)
      ? Issuer.fromPartial(object.issuerInfo)
      : undefined;
    return message;
  },
};

function createBaseQueryIdentityRequest(): QueryIdentityRequest {
  return { id: 0 };
}

export const QueryIdentityRequest = {
  encode(message: QueryIdentityRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.id !== 0) {
      writer.uint32(8).uint64(message.id);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryIdentityRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryIdentityRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.id = longToNumber(reader.uint64() as Long);
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryIdentityRequest {
    return { id: isSet(object.id) ? Number(object.id) : 0 };
  },

  toJSON(message: QueryIdentityRequest): unknown {
    const obj: any = {};
    message.id !== undefined && (obj.id = Math.round(message.id));
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryIdentityRequest>, I>>(object: I): QueryIdentityRequest {
    const message = createBaseQueryIdentityRequest();
    message.id = object.id ?? 0;
    return message;
  },
};

function createBaseQueryIdentityResponse(): QueryIdentityResponse {
  return { certificate: undefined };
}

export const QueryIdentityResponse = {
  encode(message: QueryIdentityResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.certificate !== undefined) {
      Certificate.encode(message.certificate, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryIdentityResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryIdentityResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.certificate = Certificate.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryIdentityResponse {
    return { certificate: isSet(object.certificate) ? Certificate.fromJSON(object.certificate) : undefined };
  },

  toJSON(message: QueryIdentityResponse): unknown {
    const obj: any = {};
    message.certificate !== undefined
      && (obj.certificate = message.certificate ? Certificate.toJSON(message.certificate) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryIdentityResponse>, I>>(object: I): QueryIdentityResponse {
    const message = createBaseQueryIdentityResponse();
    message.certificate = (object.certificate !== undefined && object.certificate !== null)
      ? Certificate.fromPartial(object.certificate)
      : undefined;
    return message;
  },
};

function createBaseQueryOperatorsRequest(): QueryOperatorsRequest {
  return { id: 0, pagination: undefined };
}

export const QueryOperatorsRequest = {
  encode(message: QueryOperatorsRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.id !== 0) {
      writer.uint32(8).uint64(message.id);
    }
    if (message.pagination !== undefined) {
      PageRequest.encode(message.pagination, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryOperatorsRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryOperatorsRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.id = longToNumber(reader.uint64() as Long);
          break;
        case 2:
          message.pagination = PageRequest.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryOperatorsRequest {
    return {
      id: isSet(object.id) ? Number(object.id) : 0,
      pagination: isSet(object.pagination) ? PageRequest.fromJSON(object.pagination) : undefined,
    };
  },

  toJSON(message: QueryOperatorsRequest): unknown {
    const obj: any = {};
    message.id !== undefined && (obj.id = Math.round(message.id));
    message.pagination !== undefined
      && (obj.pagination = message.pagination ? PageRequest.toJSON(message.pagination) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryOperatorsRequest>, I>>(object: I): QueryOperatorsRequest {
    const message = createBaseQueryOperatorsRequest();
    message.id = object.id ?? 0;
    message.pagination = (object.pagination !== undefined && object.pagination !== null)
      ? PageRequest.fromPartial(object.pagination)
      : undefined;
    return message;
  },
};

function createBaseQueryOperatorsResponse(): QueryOperatorsResponse {
  return { pagination: undefined };
}

export const QueryOperatorsResponse = {
  encode(message: QueryOperatorsResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.pagination !== undefined) {
      PageResponse.encode(message.pagination, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryOperatorsResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryOperatorsResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.pagination = PageResponse.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryOperatorsResponse {
    return { pagination: isSet(object.pagination) ? PageResponse.fromJSON(object.pagination) : undefined };
  },

  toJSON(message: QueryOperatorsResponse): unknown {
    const obj: any = {};
    message.pagination !== undefined
      && (obj.pagination = message.pagination ? PageResponse.toJSON(message.pagination) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryOperatorsResponse>, I>>(object: I): QueryOperatorsResponse {
    const message = createBaseQueryOperatorsResponse();
    message.pagination = (object.pagination !== undefined && object.pagination !== null)
      ? PageResponse.fromPartial(object.pagination)
      : undefined;
    return message;
  },
};

/** Query defines the gRPC querier service. */
export interface Query {
  /** Parameters queries the parameters of the module. */
  Params(request: QueryParamsRequest): Promise<QueryParamsResponse>;
  /** Queries a list of IdentityMembers items. */
  IdentityMembers(request: QueryIdentityMembersRequest): Promise<QueryIdentityMembersResponse>;
  /** Queries a list of Issuers items. */
  Issuers(request: QueryIssuersRequest): Promise<QueryIssuersResponse>;
  /** Queries a list of IssuerInfo items. */
  IssuerInfo(request: QueryIssuerInfoRequest): Promise<QueryIssuerInfoResponse>;
  /** Queries a list of Identity items. */
  Identity(request: QueryIdentityRequest): Promise<QueryIdentityResponse>;
  /** Queries a list of Operators items. */
  Operators(request: QueryOperatorsRequest): Promise<QueryOperatorsResponse>;
}

export class QueryClientImpl implements Query {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
    this.Params = this.Params.bind(this);
    this.IdentityMembers = this.IdentityMembers.bind(this);
    this.Issuers = this.Issuers.bind(this);
    this.IssuerInfo = this.IssuerInfo.bind(this);
    this.Identity = this.Identity.bind(this);
    this.Operators = this.Operators.bind(this);
  }
  Params(request: QueryParamsRequest): Promise<QueryParamsResponse> {
    const data = QueryParamsRequest.encode(request).finish();
    const promise = this.rpc.request("archive.identity.Query", "Params", data);
    return promise.then((data) => QueryParamsResponse.decode(new _m0.Reader(data)));
  }

  IdentityMembers(request: QueryIdentityMembersRequest): Promise<QueryIdentityMembersResponse> {
    const data = QueryIdentityMembersRequest.encode(request).finish();
    const promise = this.rpc.request("archive.identity.Query", "IdentityMembers", data);
    return promise.then((data) => QueryIdentityMembersResponse.decode(new _m0.Reader(data)));
  }

  Issuers(request: QueryIssuersRequest): Promise<QueryIssuersResponse> {
    const data = QueryIssuersRequest.encode(request).finish();
    const promise = this.rpc.request("archive.identity.Query", "Issuers", data);
    return promise.then((data) => QueryIssuersResponse.decode(new _m0.Reader(data)));
  }

  IssuerInfo(request: QueryIssuerInfoRequest): Promise<QueryIssuerInfoResponse> {
    const data = QueryIssuerInfoRequest.encode(request).finish();
    const promise = this.rpc.request("archive.identity.Query", "IssuerInfo", data);
    return promise.then((data) => QueryIssuerInfoResponse.decode(new _m0.Reader(data)));
  }

  Identity(request: QueryIdentityRequest): Promise<QueryIdentityResponse> {
    const data = QueryIdentityRequest.encode(request).finish();
    const promise = this.rpc.request("archive.identity.Query", "Identity", data);
    return promise.then((data) => QueryIdentityResponse.decode(new _m0.Reader(data)));
  }

  Operators(request: QueryOperatorsRequest): Promise<QueryOperatorsResponse> {
    const data = QueryOperatorsRequest.encode(request).finish();
    const promise = this.rpc.request("archive.identity.Query", "Operators", data);
    return promise.then((data) => QueryOperatorsResponse.decode(new _m0.Reader(data)));
  }
}

interface Rpc {
  request(service: string, method: string, data: Uint8Array): Promise<Uint8Array>;
}

declare var self: any | undefined;
declare var window: any | undefined;
declare var global: any | undefined;
var globalThis: any = (() => {
  if (typeof globalThis !== "undefined") {
    return globalThis;
  }
  if (typeof self !== "undefined") {
    return self;
  }
  if (typeof window !== "undefined") {
    return window;
  }
  if (typeof global !== "undefined") {
    return global;
  }
  throw "Unable to locate global object";
})();

type Builtin = Date | Function | Uint8Array | string | number | boolean | undefined;

export type DeepPartial<T> = T extends Builtin ? T
  : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>>
  : T extends {} ? { [K in keyof T]?: DeepPartial<T[K]> }
  : Partial<T>;

type KeysOfUnion<T> = T extends T ? keyof T : never;
export type Exact<P, I extends P> = P extends Builtin ? P
  : P & { [K in keyof P]: Exact<P[K], I[K]> } & { [K in Exclude<keyof I, KeysOfUnion<P>>]: never };

function longToNumber(long: Long): number {
  if (long.gt(Number.MAX_SAFE_INTEGER)) {
    throw new globalThis.Error("Value is larger than Number.MAX_SAFE_INTEGER");
  }
  return long.toNumber();
}

if (_m0.util.Long !== Long) {
  _m0.util.Long = Long as any;
  _m0.configure();
}

function isSet(value: any): boolean {
  return value !== null && value !== undefined;
}
