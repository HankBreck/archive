/* eslint-disable */
import Long from "long";
import _m0 from "protobufjs/minimal";
import { HashEntry } from "./certificate";

export const protobufPackage = "archive.identity";

export interface MsgRegisterIssuer {
  creator: string;
  name: string;
  moreInfoUri: string;
  cost: number;
}

export interface MsgRegisterIssuerResponse {
}

export interface MsgIssueCertificate {
  creator: string;
  recipient: string;
  salt: string;
  metadataSchemaUri: string;
  hashes: HashEntry[];
}

export interface MsgIssueCertificateResponse {
  id: number;
}

export interface MsgAcceptIdentity {
  creator: string;
  id: number;
}

export interface MsgAcceptIdentityResponse {
}

export interface MsgRejectIdentity {
  creator: string;
  id: number;
}

export interface MsgRejectIdentityResponse {
}

export interface MsgRevokeIdentity {
  creator: string;
  id: number;
  member: string;
}

export interface MsgRevokeIdentityResponse {
}

function createBaseMsgRegisterIssuer(): MsgRegisterIssuer {
  return { creator: "", name: "", moreInfoUri: "", cost: 0 };
}

export const MsgRegisterIssuer = {
  encode(message: MsgRegisterIssuer, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.name !== "") {
      writer.uint32(18).string(message.name);
    }
    if (message.moreInfoUri !== "") {
      writer.uint32(26).string(message.moreInfoUri);
    }
    if (message.cost !== 0) {
      writer.uint32(32).uint64(message.cost);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgRegisterIssuer {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgRegisterIssuer();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.name = reader.string();
          break;
        case 3:
          message.moreInfoUri = reader.string();
          break;
        case 4:
          message.cost = longToNumber(reader.uint64() as Long);
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgRegisterIssuer {
    return {
      creator: isSet(object.creator) ? String(object.creator) : "",
      name: isSet(object.name) ? String(object.name) : "",
      moreInfoUri: isSet(object.moreInfoUri) ? String(object.moreInfoUri) : "",
      cost: isSet(object.cost) ? Number(object.cost) : 0,
    };
  },

  toJSON(message: MsgRegisterIssuer): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.name !== undefined && (obj.name = message.name);
    message.moreInfoUri !== undefined && (obj.moreInfoUri = message.moreInfoUri);
    message.cost !== undefined && (obj.cost = Math.round(message.cost));
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgRegisterIssuer>, I>>(object: I): MsgRegisterIssuer {
    const message = createBaseMsgRegisterIssuer();
    message.creator = object.creator ?? "";
    message.name = object.name ?? "";
    message.moreInfoUri = object.moreInfoUri ?? "";
    message.cost = object.cost ?? 0;
    return message;
  },
};

function createBaseMsgRegisterIssuerResponse(): MsgRegisterIssuerResponse {
  return {};
}

export const MsgRegisterIssuerResponse = {
  encode(_: MsgRegisterIssuerResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgRegisterIssuerResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgRegisterIssuerResponse();
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

  fromJSON(_: any): MsgRegisterIssuerResponse {
    return {};
  },

  toJSON(_: MsgRegisterIssuerResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgRegisterIssuerResponse>, I>>(_: I): MsgRegisterIssuerResponse {
    const message = createBaseMsgRegisterIssuerResponse();
    return message;
  },
};

function createBaseMsgIssueCertificate(): MsgIssueCertificate {
  return { creator: "", recipient: "", salt: "", metadataSchemaUri: "", hashes: [] };
}

export const MsgIssueCertificate = {
  encode(message: MsgIssueCertificate, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.recipient !== "") {
      writer.uint32(18).string(message.recipient);
    }
    if (message.salt !== "") {
      writer.uint32(26).string(message.salt);
    }
    if (message.metadataSchemaUri !== "") {
      writer.uint32(34).string(message.metadataSchemaUri);
    }
    for (const v of message.hashes) {
      HashEntry.encode(v!, writer.uint32(42).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgIssueCertificate {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgIssueCertificate();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.recipient = reader.string();
          break;
        case 3:
          message.salt = reader.string();
          break;
        case 4:
          message.metadataSchemaUri = reader.string();
          break;
        case 5:
          message.hashes.push(HashEntry.decode(reader, reader.uint32()));
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgIssueCertificate {
    return {
      creator: isSet(object.creator) ? String(object.creator) : "",
      recipient: isSet(object.recipient) ? String(object.recipient) : "",
      salt: isSet(object.salt) ? String(object.salt) : "",
      metadataSchemaUri: isSet(object.metadataSchemaUri) ? String(object.metadataSchemaUri) : "",
      hashes: Array.isArray(object?.hashes) ? object.hashes.map((e: any) => HashEntry.fromJSON(e)) : [],
    };
  },

  toJSON(message: MsgIssueCertificate): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.recipient !== undefined && (obj.recipient = message.recipient);
    message.salt !== undefined && (obj.salt = message.salt);
    message.metadataSchemaUri !== undefined && (obj.metadataSchemaUri = message.metadataSchemaUri);
    if (message.hashes) {
      obj.hashes = message.hashes.map((e) => e ? HashEntry.toJSON(e) : undefined);
    } else {
      obj.hashes = [];
    }
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgIssueCertificate>, I>>(object: I): MsgIssueCertificate {
    const message = createBaseMsgIssueCertificate();
    message.creator = object.creator ?? "";
    message.recipient = object.recipient ?? "";
    message.salt = object.salt ?? "";
    message.metadataSchemaUri = object.metadataSchemaUri ?? "";
    message.hashes = object.hashes?.map((e) => HashEntry.fromPartial(e)) || [];
    return message;
  },
};

function createBaseMsgIssueCertificateResponse(): MsgIssueCertificateResponse {
  return { id: 0 };
}

export const MsgIssueCertificateResponse = {
  encode(message: MsgIssueCertificateResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.id !== 0) {
      writer.uint32(8).uint64(message.id);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgIssueCertificateResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgIssueCertificateResponse();
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

  fromJSON(object: any): MsgIssueCertificateResponse {
    return { id: isSet(object.id) ? Number(object.id) : 0 };
  },

  toJSON(message: MsgIssueCertificateResponse): unknown {
    const obj: any = {};
    message.id !== undefined && (obj.id = Math.round(message.id));
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgIssueCertificateResponse>, I>>(object: I): MsgIssueCertificateResponse {
    const message = createBaseMsgIssueCertificateResponse();
    message.id = object.id ?? 0;
    return message;
  },
};

function createBaseMsgAcceptIdentity(): MsgAcceptIdentity {
  return { creator: "", id: 0 };
}

export const MsgAcceptIdentity = {
  encode(message: MsgAcceptIdentity, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.id !== 0) {
      writer.uint32(16).uint64(message.id);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgAcceptIdentity {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgAcceptIdentity();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.id = longToNumber(reader.uint64() as Long);
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgAcceptIdentity {
    return {
      creator: isSet(object.creator) ? String(object.creator) : "",
      id: isSet(object.id) ? Number(object.id) : 0,
    };
  },

  toJSON(message: MsgAcceptIdentity): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.id !== undefined && (obj.id = Math.round(message.id));
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgAcceptIdentity>, I>>(object: I): MsgAcceptIdentity {
    const message = createBaseMsgAcceptIdentity();
    message.creator = object.creator ?? "";
    message.id = object.id ?? 0;
    return message;
  },
};

function createBaseMsgAcceptIdentityResponse(): MsgAcceptIdentityResponse {
  return {};
}

export const MsgAcceptIdentityResponse = {
  encode(_: MsgAcceptIdentityResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgAcceptIdentityResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgAcceptIdentityResponse();
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

  fromJSON(_: any): MsgAcceptIdentityResponse {
    return {};
  },

  toJSON(_: MsgAcceptIdentityResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgAcceptIdentityResponse>, I>>(_: I): MsgAcceptIdentityResponse {
    const message = createBaseMsgAcceptIdentityResponse();
    return message;
  },
};

function createBaseMsgRejectIdentity(): MsgRejectIdentity {
  return { creator: "", id: 0 };
}

export const MsgRejectIdentity = {
  encode(message: MsgRejectIdentity, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.id !== 0) {
      writer.uint32(16).uint64(message.id);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgRejectIdentity {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgRejectIdentity();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.id = longToNumber(reader.uint64() as Long);
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgRejectIdentity {
    return {
      creator: isSet(object.creator) ? String(object.creator) : "",
      id: isSet(object.id) ? Number(object.id) : 0,
    };
  },

  toJSON(message: MsgRejectIdentity): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.id !== undefined && (obj.id = Math.round(message.id));
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgRejectIdentity>, I>>(object: I): MsgRejectIdentity {
    const message = createBaseMsgRejectIdentity();
    message.creator = object.creator ?? "";
    message.id = object.id ?? 0;
    return message;
  },
};

function createBaseMsgRejectIdentityResponse(): MsgRejectIdentityResponse {
  return {};
}

export const MsgRejectIdentityResponse = {
  encode(_: MsgRejectIdentityResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgRejectIdentityResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgRejectIdentityResponse();
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

  fromJSON(_: any): MsgRejectIdentityResponse {
    return {};
  },

  toJSON(_: MsgRejectIdentityResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgRejectIdentityResponse>, I>>(_: I): MsgRejectIdentityResponse {
    const message = createBaseMsgRejectIdentityResponse();
    return message;
  },
};

function createBaseMsgRevokeIdentity(): MsgRevokeIdentity {
  return { creator: "", id: 0, member: "" };
}

export const MsgRevokeIdentity = {
  encode(message: MsgRevokeIdentity, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.id !== 0) {
      writer.uint32(16).uint64(message.id);
    }
    if (message.member !== "") {
      writer.uint32(26).string(message.member);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgRevokeIdentity {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgRevokeIdentity();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.id = longToNumber(reader.uint64() as Long);
          break;
        case 3:
          message.member = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgRevokeIdentity {
    return {
      creator: isSet(object.creator) ? String(object.creator) : "",
      id: isSet(object.id) ? Number(object.id) : 0,
      member: isSet(object.member) ? String(object.member) : "",
    };
  },

  toJSON(message: MsgRevokeIdentity): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.id !== undefined && (obj.id = Math.round(message.id));
    message.member !== undefined && (obj.member = message.member);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgRevokeIdentity>, I>>(object: I): MsgRevokeIdentity {
    const message = createBaseMsgRevokeIdentity();
    message.creator = object.creator ?? "";
    message.id = object.id ?? 0;
    message.member = object.member ?? "";
    return message;
  },
};

function createBaseMsgRevokeIdentityResponse(): MsgRevokeIdentityResponse {
  return {};
}

export const MsgRevokeIdentityResponse = {
  encode(_: MsgRevokeIdentityResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgRevokeIdentityResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgRevokeIdentityResponse();
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

  fromJSON(_: any): MsgRevokeIdentityResponse {
    return {};
  },

  toJSON(_: MsgRevokeIdentityResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgRevokeIdentityResponse>, I>>(_: I): MsgRevokeIdentityResponse {
    const message = createBaseMsgRevokeIdentityResponse();
    return message;
  },
};

/** Msg defines the Msg service. */
export interface Msg {
  RegisterIssuer(request: MsgRegisterIssuer): Promise<MsgRegisterIssuerResponse>;
  IssueCertificate(request: MsgIssueCertificate): Promise<MsgIssueCertificateResponse>;
  AcceptIdentity(request: MsgAcceptIdentity): Promise<MsgAcceptIdentityResponse>;
  RejectIdentity(request: MsgRejectIdentity): Promise<MsgRejectIdentityResponse>;
  /** this line is used by starport scaffolding # proto/tx/rpc */
  RevokeIdentity(request: MsgRevokeIdentity): Promise<MsgRevokeIdentityResponse>;
}

export class MsgClientImpl implements Msg {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
    this.RegisterIssuer = this.RegisterIssuer.bind(this);
    this.IssueCertificate = this.IssueCertificate.bind(this);
    this.AcceptIdentity = this.AcceptIdentity.bind(this);
    this.RejectIdentity = this.RejectIdentity.bind(this);
    this.RevokeIdentity = this.RevokeIdentity.bind(this);
  }
  RegisterIssuer(request: MsgRegisterIssuer): Promise<MsgRegisterIssuerResponse> {
    const data = MsgRegisterIssuer.encode(request).finish();
    const promise = this.rpc.request("archive.identity.Msg", "RegisterIssuer", data);
    return promise.then((data) => MsgRegisterIssuerResponse.decode(new _m0.Reader(data)));
  }

  IssueCertificate(request: MsgIssueCertificate): Promise<MsgIssueCertificateResponse> {
    const data = MsgIssueCertificate.encode(request).finish();
    const promise = this.rpc.request("archive.identity.Msg", "IssueCertificate", data);
    return promise.then((data) => MsgIssueCertificateResponse.decode(new _m0.Reader(data)));
  }

  AcceptIdentity(request: MsgAcceptIdentity): Promise<MsgAcceptIdentityResponse> {
    const data = MsgAcceptIdentity.encode(request).finish();
    const promise = this.rpc.request("archive.identity.Msg", "AcceptIdentity", data);
    return promise.then((data) => MsgAcceptIdentityResponse.decode(new _m0.Reader(data)));
  }

  RejectIdentity(request: MsgRejectIdentity): Promise<MsgRejectIdentityResponse> {
    const data = MsgRejectIdentity.encode(request).finish();
    const promise = this.rpc.request("archive.identity.Msg", "RejectIdentity", data);
    return promise.then((data) => MsgRejectIdentityResponse.decode(new _m0.Reader(data)));
  }

  RevokeIdentity(request: MsgRevokeIdentity): Promise<MsgRevokeIdentityResponse> {
    const data = MsgRevokeIdentity.encode(request).finish();
    const promise = this.rpc.request("archive.identity.Msg", "RevokeIdentity", data);
    return promise.then((data) => MsgRevokeIdentityResponse.decode(new _m0.Reader(data)));
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
