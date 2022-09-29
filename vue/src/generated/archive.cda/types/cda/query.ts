/* eslint-disable */
import { Reader, util, configure, Writer } from "protobufjs/minimal";
import * as Long from "long";
import { Params } from "../cda/params";
import { CDA } from "../cda/cda";
import {
  PageRequest,
  PageResponse,
} from "../cosmos/base/query/v1beta1/pagination";

export const protobufPackage = "archive.cda";

/** QueryParamsRequest is request type for the Query/Params RPC method. */
export interface QueryParamsRequest {}

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

export interface QueryApprovalsRequest {
  cdaId: string;
}

export interface QueryApprovalsResponse {
  approvals: string;
}

const baseQueryParamsRequest: object = {};

export const QueryParamsRequest = {
  encode(_: QueryParamsRequest, writer: Writer = Writer.create()): Writer {
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryParamsRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryParamsRequest } as QueryParamsRequest;
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
    const message = { ...baseQueryParamsRequest } as QueryParamsRequest;
    return message;
  },

  toJSON(_: QueryParamsRequest): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(_: DeepPartial<QueryParamsRequest>): QueryParamsRequest {
    const message = { ...baseQueryParamsRequest } as QueryParamsRequest;
    return message;
  },
};

const baseQueryParamsResponse: object = {};

export const QueryParamsResponse = {
  encode(
    message: QueryParamsResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.params !== undefined) {
      Params.encode(message.params, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryParamsResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryParamsResponse } as QueryParamsResponse;
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
    const message = { ...baseQueryParamsResponse } as QueryParamsResponse;
    if (object.params !== undefined && object.params !== null) {
      message.params = Params.fromJSON(object.params);
    } else {
      message.params = undefined;
    }
    return message;
  },

  toJSON(message: QueryParamsResponse): unknown {
    const obj: any = {};
    message.params !== undefined &&
      (obj.params = message.params ? Params.toJSON(message.params) : undefined);
    return obj;
  },

  fromPartial(object: DeepPartial<QueryParamsResponse>): QueryParamsResponse {
    const message = { ...baseQueryParamsResponse } as QueryParamsResponse;
    if (object.params !== undefined && object.params !== null) {
      message.params = Params.fromPartial(object.params);
    } else {
      message.params = undefined;
    }
    return message;
  },
};

const baseQueryCdaRequest: object = { id: 0 };

export const QueryCdaRequest = {
  encode(message: QueryCdaRequest, writer: Writer = Writer.create()): Writer {
    if (message.id !== 0) {
      writer.uint32(8).uint64(message.id);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryCdaRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryCdaRequest } as QueryCdaRequest;
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
    const message = { ...baseQueryCdaRequest } as QueryCdaRequest;
    if (object.id !== undefined && object.id !== null) {
      message.id = Number(object.id);
    } else {
      message.id = 0;
    }
    return message;
  },

  toJSON(message: QueryCdaRequest): unknown {
    const obj: any = {};
    message.id !== undefined && (obj.id = message.id);
    return obj;
  },

  fromPartial(object: DeepPartial<QueryCdaRequest>): QueryCdaRequest {
    const message = { ...baseQueryCdaRequest } as QueryCdaRequest;
    if (object.id !== undefined && object.id !== null) {
      message.id = object.id;
    } else {
      message.id = 0;
    }
    return message;
  },
};

const baseQueryCdaResponse: object = {};

export const QueryCdaResponse = {
  encode(message: QueryCdaResponse, writer: Writer = Writer.create()): Writer {
    if (message.cda !== undefined) {
      CDA.encode(message.cda, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryCdaResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryCdaResponse } as QueryCdaResponse;
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
    const message = { ...baseQueryCdaResponse } as QueryCdaResponse;
    if (object.cda !== undefined && object.cda !== null) {
      message.cda = CDA.fromJSON(object.cda);
    } else {
      message.cda = undefined;
    }
    return message;
  },

  toJSON(message: QueryCdaResponse): unknown {
    const obj: any = {};
    message.cda !== undefined &&
      (obj.cda = message.cda ? CDA.toJSON(message.cda) : undefined);
    return obj;
  },

  fromPartial(object: DeepPartial<QueryCdaResponse>): QueryCdaResponse {
    const message = { ...baseQueryCdaResponse } as QueryCdaResponse;
    if (object.cda !== undefined && object.cda !== null) {
      message.cda = CDA.fromPartial(object.cda);
    } else {
      message.cda = undefined;
    }
    return message;
  },
};

const baseQueryCdasRequest: object = {};

export const QueryCdasRequest = {
  encode(message: QueryCdasRequest, writer: Writer = Writer.create()): Writer {
    if (message.pagination !== undefined) {
      PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryCdasRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryCdasRequest } as QueryCdasRequest;
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
    const message = { ...baseQueryCdasRequest } as QueryCdasRequest;
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryCdasRequest): unknown {
    const obj: any = {};
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageRequest.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(object: DeepPartial<QueryCdasRequest>): QueryCdasRequest {
    const message = { ...baseQueryCdasRequest } as QueryCdasRequest;
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromPartial(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },
};

const baseQueryCdasResponse: object = {};

export const QueryCdasResponse = {
  encode(message: QueryCdasResponse, writer: Writer = Writer.create()): Writer {
    for (const v of message.CDAs) {
      CDA.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    if (message.pagination !== undefined) {
      PageResponse.encode(
        message.pagination,
        writer.uint32(18).fork()
      ).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryCdasResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryCdasResponse } as QueryCdasResponse;
    message.CDAs = [];
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
    const message = { ...baseQueryCdasResponse } as QueryCdasResponse;
    message.CDAs = [];
    if (object.CDAs !== undefined && object.CDAs !== null) {
      for (const e of object.CDAs) {
        message.CDAs.push(CDA.fromJSON(e));
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryCdasResponse): unknown {
    const obj: any = {};
    if (message.CDAs) {
      obj.CDAs = message.CDAs.map((e) => (e ? CDA.toJSON(e) : undefined));
    } else {
      obj.CDAs = [];
    }
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageResponse.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(object: DeepPartial<QueryCdasResponse>): QueryCdasResponse {
    const message = { ...baseQueryCdasResponse } as QueryCdasResponse;
    message.CDAs = [];
    if (object.CDAs !== undefined && object.CDAs !== null) {
      for (const e of object.CDAs) {
        message.CDAs.push(CDA.fromPartial(e));
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromPartial(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },
};

const baseQueryCdasOwnedRequest: object = { owner: "" };

export const QueryCdasOwnedRequest = {
  encode(
    message: QueryCdasOwnedRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.owner !== "") {
      writer.uint32(10).string(message.owner);
    }
    if (message.pagination !== undefined) {
      PageRequest.encode(message.pagination, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryCdasOwnedRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryCdasOwnedRequest } as QueryCdasOwnedRequest;
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
    const message = { ...baseQueryCdasOwnedRequest } as QueryCdasOwnedRequest;
    if (object.owner !== undefined && object.owner !== null) {
      message.owner = String(object.owner);
    } else {
      message.owner = "";
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryCdasOwnedRequest): unknown {
    const obj: any = {};
    message.owner !== undefined && (obj.owner = message.owner);
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageRequest.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryCdasOwnedRequest>
  ): QueryCdasOwnedRequest {
    const message = { ...baseQueryCdasOwnedRequest } as QueryCdasOwnedRequest;
    if (object.owner !== undefined && object.owner !== null) {
      message.owner = object.owner;
    } else {
      message.owner = "";
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromPartial(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },
};

const baseQueryCdasOwnedResponse: object = { ids: 0 };

export const QueryCdasOwnedResponse = {
  encode(
    message: QueryCdasOwnedResponse,
    writer: Writer = Writer.create()
  ): Writer {
    writer.uint32(10).fork();
    for (const v of message.ids) {
      writer.uint64(v);
    }
    writer.ldelim();
    if (message.pagination !== undefined) {
      PageResponse.encode(
        message.pagination,
        writer.uint32(18).fork()
      ).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryCdasOwnedResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryCdasOwnedResponse } as QueryCdasOwnedResponse;
    message.ids = [];
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
    const message = { ...baseQueryCdasOwnedResponse } as QueryCdasOwnedResponse;
    message.ids = [];
    if (object.ids !== undefined && object.ids !== null) {
      for (const e of object.ids) {
        message.ids.push(Number(e));
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryCdasOwnedResponse): unknown {
    const obj: any = {};
    if (message.ids) {
      obj.ids = message.ids.map((e) => e);
    } else {
      obj.ids = [];
    }
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageResponse.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryCdasOwnedResponse>
  ): QueryCdasOwnedResponse {
    const message = { ...baseQueryCdasOwnedResponse } as QueryCdasOwnedResponse;
    message.ids = [];
    if (object.ids !== undefined && object.ids !== null) {
      for (const e of object.ids) {
        message.ids.push(e);
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromPartial(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },
};

const baseQueryApprovalsRequest: object = { cdaId: "" };

export const QueryApprovalsRequest = {
  encode(
    message: QueryApprovalsRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.cdaId !== "") {
      writer.uint32(10).string(message.cdaId);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryApprovalsRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryApprovalsRequest } as QueryApprovalsRequest;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.cdaId = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryApprovalsRequest {
    const message = { ...baseQueryApprovalsRequest } as QueryApprovalsRequest;
    if (object.cdaId !== undefined && object.cdaId !== null) {
      message.cdaId = String(object.cdaId);
    } else {
      message.cdaId = "";
    }
    return message;
  },

  toJSON(message: QueryApprovalsRequest): unknown {
    const obj: any = {};
    message.cdaId !== undefined && (obj.cdaId = message.cdaId);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryApprovalsRequest>
  ): QueryApprovalsRequest {
    const message = { ...baseQueryApprovalsRequest } as QueryApprovalsRequest;
    if (object.cdaId !== undefined && object.cdaId !== null) {
      message.cdaId = object.cdaId;
    } else {
      message.cdaId = "";
    }
    return message;
  },
};

const baseQueryApprovalsResponse: object = { approvals: "" };

export const QueryApprovalsResponse = {
  encode(
    message: QueryApprovalsResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.approvals !== "") {
      writer.uint32(10).string(message.approvals);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryApprovalsResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryApprovalsResponse } as QueryApprovalsResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.approvals = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryApprovalsResponse {
    const message = { ...baseQueryApprovalsResponse } as QueryApprovalsResponse;
    if (object.approvals !== undefined && object.approvals !== null) {
      message.approvals = String(object.approvals);
    } else {
      message.approvals = "";
    }
    return message;
  },

  toJSON(message: QueryApprovalsResponse): unknown {
    const obj: any = {};
    message.approvals !== undefined && (obj.approvals = message.approvals);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryApprovalsResponse>
  ): QueryApprovalsResponse {
    const message = { ...baseQueryApprovalsResponse } as QueryApprovalsResponse;
    if (object.approvals !== undefined && object.approvals !== null) {
      message.approvals = object.approvals;
    } else {
      message.approvals = "";
    }
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
  Approvals(request: QueryApprovalsRequest): Promise<QueryApprovalsResponse>;
}

export class QueryClientImpl implements Query {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
  }
  Params(request: QueryParamsRequest): Promise<QueryParamsResponse> {
    const data = QueryParamsRequest.encode(request).finish();
    const promise = this.rpc.request("archive.cda.Query", "Params", data);
    return promise.then((data) => QueryParamsResponse.decode(new Reader(data)));
  }

  Cda(request: QueryCdaRequest): Promise<QueryCdaResponse> {
    const data = QueryCdaRequest.encode(request).finish();
    const promise = this.rpc.request("archive.cda.Query", "Cda", data);
    return promise.then((data) => QueryCdaResponse.decode(new Reader(data)));
  }

  Cdas(request: QueryCdasRequest): Promise<QueryCdasResponse> {
    const data = QueryCdasRequest.encode(request).finish();
    const promise = this.rpc.request("archive.cda.Query", "Cdas", data);
    return promise.then((data) => QueryCdasResponse.decode(new Reader(data)));
  }

  CdasOwned(request: QueryCdasOwnedRequest): Promise<QueryCdasOwnedResponse> {
    const data = QueryCdasOwnedRequest.encode(request).finish();
    const promise = this.rpc.request("archive.cda.Query", "CdasOwned", data);
    return promise.then((data) =>
      QueryCdasOwnedResponse.decode(new Reader(data))
    );
  }

  Approvals(request: QueryApprovalsRequest): Promise<QueryApprovalsResponse> {
    const data = QueryApprovalsRequest.encode(request).finish();
    const promise = this.rpc.request("archive.cda.Query", "Approvals", data);
    return promise.then((data) =>
      QueryApprovalsResponse.decode(new Reader(data))
    );
  }
}

interface Rpc {
  request(
    service: string,
    method: string,
    data: Uint8Array
  ): Promise<Uint8Array>;
}

declare var self: any | undefined;
declare var window: any | undefined;
var globalThis: any = (() => {
  if (typeof globalThis !== "undefined") return globalThis;
  if (typeof self !== "undefined") return self;
  if (typeof window !== "undefined") return window;
  if (typeof global !== "undefined") return global;
  throw "Unable to locate global object";
})();

type Builtin = Date | Function | Uint8Array | string | number | undefined;
export type DeepPartial<T> = T extends Builtin
  ? T
  : T extends Array<infer U>
  ? Array<DeepPartial<U>>
  : T extends ReadonlyArray<infer U>
  ? ReadonlyArray<DeepPartial<U>>
  : T extends {}
  ? { [K in keyof T]?: DeepPartial<T[K]> }
  : Partial<T>;

function longToNumber(long: Long): number {
  if (long.gt(Number.MAX_SAFE_INTEGER)) {
    throw new globalThis.Error("Value is larger than Number.MAX_SAFE_INTEGER");
  }
  return long.toNumber();
}

if (util.Long !== Long) {
  util.Long = Long as any;
  configure();
}
