/* eslint-disable */
import Long from "long";
import _m0 from "protobufjs/minimal";
import { PageRequest, PageResponse } from "../../cosmos/base/query/v1beta1/pagination";
import { CDA } from "./cda";
import { Params } from "./params";

export const protobufPackage = "archive.cda";

/** QueryParamsRequest is request type for the Query/Params RPC method. */
export interface QueryParamsRequest {
}

/** QueryParamsResponse is response type for the Query/Params RPC method. */
export interface QueryParamsResponse {
  /** params holds all the parameters of this module. */
  params: Params | undefined;
}

export interface QueryCdaRequest {
  id: number;
}

export interface QueryCdaResponse {
  cda: CDA | undefined;
}

export interface QueryCdasRequest {
  /** Pagination to view all CDAs */
  pagination: PageRequest | undefined;
}

export interface QueryCdasResponse {
  /** List of CDA objects */
  CDAs: CDA[];
  /** Pagination to view all CDAs */
  pagination: PageResponse | undefined;
}

export interface QueryCdasOwnedRequest {
  /** Account address for the owner */
  owner: string;
  /** Pagination to view all ids */
  pagination: PageRequest | undefined;
}

export interface QueryCdasOwnedResponse {
  /** List of CDA ids belonging to the owner */
  ids: number[];
  /** Pagination to view all CDAs */
  pagination: PageResponse | undefined;
}

export interface QueryApprovalRequest {
  /** The id of the CDA to check */
  cdaId: number;
  /** The wallet address of the owner to check */
  owner: string;
}

export interface QueryApprovalResponse {
  approved: boolean;
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

function createBaseQueryCdaRequest(): QueryCdaRequest {
  return { id: 0 };
}

export const QueryCdaRequest = {
  encode(message: QueryCdaRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.id !== 0) {
      writer.uint32(8).uint64(message.id);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryCdaRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryCdaRequest();
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

  fromJSON(object: any): QueryCdaRequest {
    return { id: isSet(object.id) ? Number(object.id) : 0 };
  },

  toJSON(message: QueryCdaRequest): unknown {
    const obj: any = {};
    message.id !== undefined && (obj.id = Math.round(message.id));
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryCdaRequest>, I>>(object: I): QueryCdaRequest {
    const message = createBaseQueryCdaRequest();
    message.id = object.id ?? 0;
    return message;
  },
};

function createBaseQueryCdaResponse(): QueryCdaResponse {
  return { cda: undefined };
}

export const QueryCdaResponse = {
  encode(message: QueryCdaResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.cda !== undefined) {
      CDA.encode(message.cda, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryCdaResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryCdaResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.cda = CDA.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryCdaResponse {
    return { cda: isSet(object.cda) ? CDA.fromJSON(object.cda) : undefined };
  },

  toJSON(message: QueryCdaResponse): unknown {
    const obj: any = {};
    message.cda !== undefined && (obj.cda = message.cda ? CDA.toJSON(message.cda) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryCdaResponse>, I>>(object: I): QueryCdaResponse {
    const message = createBaseQueryCdaResponse();
    message.cda = (object.cda !== undefined && object.cda !== null) ? CDA.fromPartial(object.cda) : undefined;
    return message;
  },
};

function createBaseQueryCdasRequest(): QueryCdasRequest {
  return { pagination: undefined };
}

export const QueryCdasRequest = {
  encode(message: QueryCdasRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.pagination !== undefined) {
      PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryCdasRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryCdasRequest();
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

  fromJSON(object: any): QueryCdasRequest {
    return { pagination: isSet(object.pagination) ? PageRequest.fromJSON(object.pagination) : undefined };
  },

  toJSON(message: QueryCdasRequest): unknown {
    const obj: any = {};
    message.pagination !== undefined
      && (obj.pagination = message.pagination ? PageRequest.toJSON(message.pagination) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryCdasRequest>, I>>(object: I): QueryCdasRequest {
    const message = createBaseQueryCdasRequest();
    message.pagination = (object.pagination !== undefined && object.pagination !== null)
      ? PageRequest.fromPartial(object.pagination)
      : undefined;
    return message;
  },
};

function createBaseQueryCdasResponse(): QueryCdasResponse {
  return { CDAs: [], pagination: undefined };
}

export const QueryCdasResponse = {
  encode(message: QueryCdasResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    for (const v of message.CDAs) {
      CDA.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    if (message.pagination !== undefined) {
      PageResponse.encode(message.pagination, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryCdasResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryCdasResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.CDAs.push(CDA.decode(reader, reader.uint32()));
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

  fromJSON(object: any): QueryCdasResponse {
    return {
      CDAs: Array.isArray(object?.CDAs) ? object.CDAs.map((e: any) => CDA.fromJSON(e)) : [],
      pagination: isSet(object.pagination) ? PageResponse.fromJSON(object.pagination) : undefined,
    };
  },

  toJSON(message: QueryCdasResponse): unknown {
    const obj: any = {};
    if (message.CDAs) {
      obj.CDAs = message.CDAs.map((e) => e ? CDA.toJSON(e) : undefined);
    } else {
      obj.CDAs = [];
    }
    message.pagination !== undefined
      && (obj.pagination = message.pagination ? PageResponse.toJSON(message.pagination) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryCdasResponse>, I>>(object: I): QueryCdasResponse {
    const message = createBaseQueryCdasResponse();
    message.CDAs = object.CDAs?.map((e) => CDA.fromPartial(e)) || [];
    message.pagination = (object.pagination !== undefined && object.pagination !== null)
      ? PageResponse.fromPartial(object.pagination)
      : undefined;
    return message;
  },
};

function createBaseQueryCdasOwnedRequest(): QueryCdasOwnedRequest {
  return { owner: "", pagination: undefined };
}

export const QueryCdasOwnedRequest = {
  encode(message: QueryCdasOwnedRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.owner !== "") {
      writer.uint32(10).string(message.owner);
    }
    if (message.pagination !== undefined) {
      PageRequest.encode(message.pagination, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryCdasOwnedRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryCdasOwnedRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.owner = reader.string();
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

  fromJSON(object: any): QueryCdasOwnedRequest {
    return {
      owner: isSet(object.owner) ? String(object.owner) : "",
      pagination: isSet(object.pagination) ? PageRequest.fromJSON(object.pagination) : undefined,
    };
  },

  toJSON(message: QueryCdasOwnedRequest): unknown {
    const obj: any = {};
    message.owner !== undefined && (obj.owner = message.owner);
    message.pagination !== undefined
      && (obj.pagination = message.pagination ? PageRequest.toJSON(message.pagination) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryCdasOwnedRequest>, I>>(object: I): QueryCdasOwnedRequest {
    const message = createBaseQueryCdasOwnedRequest();
    message.owner = object.owner ?? "";
    message.pagination = (object.pagination !== undefined && object.pagination !== null)
      ? PageRequest.fromPartial(object.pagination)
      : undefined;
    return message;
  },
};

function createBaseQueryCdasOwnedResponse(): QueryCdasOwnedResponse {
  return { ids: [], pagination: undefined };
}

export const QueryCdasOwnedResponse = {
  encode(message: QueryCdasOwnedResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    writer.uint32(10).fork();
    for (const v of message.ids) {
      writer.uint64(v);
    }
    writer.ldelim();
    if (message.pagination !== undefined) {
      PageResponse.encode(message.pagination, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryCdasOwnedResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryCdasOwnedResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if ((tag & 7) === 2) {
            const end2 = reader.uint32() + reader.pos;
            while (reader.pos < end2) {
              message.ids.push(longToNumber(reader.uint64() as Long));
            }
          } else {
            message.ids.push(longToNumber(reader.uint64() as Long));
          }
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

  fromJSON(object: any): QueryCdasOwnedResponse {
    return {
      ids: Array.isArray(object?.ids) ? object.ids.map((e: any) => Number(e)) : [],
      pagination: isSet(object.pagination) ? PageResponse.fromJSON(object.pagination) : undefined,
    };
  },

  toJSON(message: QueryCdasOwnedResponse): unknown {
    const obj: any = {};
    if (message.ids) {
      obj.ids = message.ids.map((e) => Math.round(e));
    } else {
      obj.ids = [];
    }
    message.pagination !== undefined
      && (obj.pagination = message.pagination ? PageResponse.toJSON(message.pagination) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryCdasOwnedResponse>, I>>(object: I): QueryCdasOwnedResponse {
    const message = createBaseQueryCdasOwnedResponse();
    message.ids = object.ids?.map((e) => e) || [];
    message.pagination = (object.pagination !== undefined && object.pagination !== null)
      ? PageResponse.fromPartial(object.pagination)
      : undefined;
    return message;
  },
};

function createBaseQueryApprovalRequest(): QueryApprovalRequest {
  return { cdaId: 0, owner: "" };
}

export const QueryApprovalRequest = {
  encode(message: QueryApprovalRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.cdaId !== 0) {
      writer.uint32(8).uint64(message.cdaId);
    }
    if (message.owner !== "") {
      writer.uint32(18).string(message.owner);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryApprovalRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryApprovalRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.cdaId = longToNumber(reader.uint64() as Long);
          break;
        case 2:
          message.owner = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryApprovalRequest {
    return {
      cdaId: isSet(object.cdaId) ? Number(object.cdaId) : 0,
      owner: isSet(object.owner) ? String(object.owner) : "",
    };
  },

  toJSON(message: QueryApprovalRequest): unknown {
    const obj: any = {};
    message.cdaId !== undefined && (obj.cdaId = Math.round(message.cdaId));
    message.owner !== undefined && (obj.owner = message.owner);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryApprovalRequest>, I>>(object: I): QueryApprovalRequest {
    const message = createBaseQueryApprovalRequest();
    message.cdaId = object.cdaId ?? 0;
    message.owner = object.owner ?? "";
    return message;
  },
};

function createBaseQueryApprovalResponse(): QueryApprovalResponse {
  return { approved: false };
}

export const QueryApprovalResponse = {
  encode(message: QueryApprovalResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.approved === true) {
      writer.uint32(8).bool(message.approved);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryApprovalResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryApprovalResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.approved = reader.bool();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryApprovalResponse {
    return { approved: isSet(object.approved) ? Boolean(object.approved) : false };
  },

  toJSON(message: QueryApprovalResponse): unknown {
    const obj: any = {};
    message.approved !== undefined && (obj.approved = message.approved);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryApprovalResponse>, I>>(object: I): QueryApprovalResponse {
    const message = createBaseQueryApprovalResponse();
    message.approved = object.approved ?? false;
    return message;
  },
};

/** Query defines the gRPC querier service. */
export interface Query {
  /** Parameters queries the parameters of the module. */
  Params(request: QueryParamsRequest): Promise<QueryParamsResponse>;
  /** Queries a list of Cda items. */
  Cda(request: QueryCdaRequest): Promise<QueryCdaResponse>;
  /** Queries a list of Cdas items. */
  Cdas(request: QueryCdasRequest): Promise<QueryCdasResponse>;
  /** Queries a list of CdasOwned items. */
  CdasOwned(request: QueryCdasOwnedRequest): Promise<QueryCdasOwnedResponse>;
  /** Queries a list of Approvals items. */
  Approval(request: QueryApprovalRequest): Promise<QueryApprovalResponse>;
}

export class QueryClientImpl implements Query {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
    this.Params = this.Params.bind(this);
    this.Cda = this.Cda.bind(this);
    this.Cdas = this.Cdas.bind(this);
    this.CdasOwned = this.CdasOwned.bind(this);
    this.Approval = this.Approval.bind(this);
  }
  Params(request: QueryParamsRequest): Promise<QueryParamsResponse> {
    const data = QueryParamsRequest.encode(request).finish();
    const promise = this.rpc.request("archive.cda.Query", "Params", data);
    return promise.then((data) => QueryParamsResponse.decode(new _m0.Reader(data)));
  }

  Cda(request: QueryCdaRequest): Promise<QueryCdaResponse> {
    const data = QueryCdaRequest.encode(request).finish();
    const promise = this.rpc.request("archive.cda.Query", "Cda", data);
    return promise.then((data) => QueryCdaResponse.decode(new _m0.Reader(data)));
  }

  Cdas(request: QueryCdasRequest): Promise<QueryCdasResponse> {
    const data = QueryCdasRequest.encode(request).finish();
    const promise = this.rpc.request("archive.cda.Query", "Cdas", data);
    return promise.then((data) => QueryCdasResponse.decode(new _m0.Reader(data)));
  }

  CdasOwned(request: QueryCdasOwnedRequest): Promise<QueryCdasOwnedResponse> {
    const data = QueryCdasOwnedRequest.encode(request).finish();
    const promise = this.rpc.request("archive.cda.Query", "CdasOwned", data);
    return promise.then((data) => QueryCdasOwnedResponse.decode(new _m0.Reader(data)));
  }

  Approval(request: QueryApprovalRequest): Promise<QueryApprovalResponse> {
    const data = QueryApprovalRequest.encode(request).finish();
    const promise = this.rpc.request("archive.cda.Query", "Approval", data);
    return promise.then((data) => QueryApprovalResponse.decode(new _m0.Reader(data)));
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
