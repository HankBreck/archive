/* eslint-disable */
import Long from "long";
import _m0 from "protobufjs/minimal";

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
  hashes: { [key: string]: string };
}

export interface MsgIssueCertificate_HashesEntry {
  key: string;
  value: string;
}

export interface MsgIssueCertificateResponse {
  id: number;
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
  return { creator: "", recipient: "", salt: "", metadataSchemaUri: "", hashes: {} };
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
    Object.entries(message.hashes).forEach(([key, value]) => {
      MsgIssueCertificate_HashesEntry.encode({ key: key as any, value }, writer.uint32(42).fork()).ldelim();
    });
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
          const entry5 = MsgIssueCertificate_HashesEntry.decode(reader, reader.uint32());
          if (entry5.value !== undefined) {
            message.hashes[entry5.key] = entry5.value;
          }
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
      hashes: isObject(object.hashes)
        ? Object.entries(object.hashes).reduce<{ [key: string]: string }>((acc, [key, value]) => {
          acc[key] = String(value);
          return acc;
        }, {})
        : {},
    };
  },

  toJSON(message: MsgIssueCertificate): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.recipient !== undefined && (obj.recipient = message.recipient);
    message.salt !== undefined && (obj.salt = message.salt);
    message.metadataSchemaUri !== undefined && (obj.metadataSchemaUri = message.metadataSchemaUri);
    obj.hashes = {};
    if (message.hashes) {
      Object.entries(message.hashes).forEach(([k, v]) => {
        obj.hashes[k] = v;
      });
    }
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgIssueCertificate>, I>>(object: I): MsgIssueCertificate {
    const message = createBaseMsgIssueCertificate();
    message.creator = object.creator ?? "";
    message.recipient = object.recipient ?? "";
    message.salt = object.salt ?? "";
    message.metadataSchemaUri = object.metadataSchemaUri ?? "";
    message.hashes = Object.entries(object.hashes ?? {}).reduce<{ [key: string]: string }>((acc, [key, value]) => {
      if (value !== undefined) {
        acc[key] = String(value);
      }
      return acc;
    }, {});
    return message;
  },
};

function createBaseMsgIssueCertificate_HashesEntry(): MsgIssueCertificate_HashesEntry {
  return { key: "", value: "" };
}

export const MsgIssueCertificate_HashesEntry = {
  encode(message: MsgIssueCertificate_HashesEntry, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.key !== "") {
      writer.uint32(10).string(message.key);
    }
    if (message.value !== "") {
      writer.uint32(18).string(message.value);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgIssueCertificate_HashesEntry {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgIssueCertificate_HashesEntry();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.key = reader.string();
          break;
        case 2:
          message.value = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgIssueCertificate_HashesEntry {
    return { key: isSet(object.key) ? String(object.key) : "", value: isSet(object.value) ? String(object.value) : "" };
  },

  toJSON(message: MsgIssueCertificate_HashesEntry): unknown {
    const obj: any = {};
    message.key !== undefined && (obj.key = message.key);
    message.value !== undefined && (obj.value = message.value);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgIssueCertificate_HashesEntry>, I>>(
    object: I,
  ): MsgIssueCertificate_HashesEntry {
    const message = createBaseMsgIssueCertificate_HashesEntry();
    message.key = object.key ?? "";
    message.value = object.value ?? "";
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

/** Msg defines the Msg service. */
export interface Msg {
  RegisterIssuer(request: MsgRegisterIssuer): Promise<MsgRegisterIssuerResponse>;
  /** this line is used by starport scaffolding # proto/tx/rpc */
  IssueCertificate(request: MsgIssueCertificate): Promise<MsgIssueCertificateResponse>;
}

export class MsgClientImpl implements Msg {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
    this.RegisterIssuer = this.RegisterIssuer.bind(this);
    this.IssueCertificate = this.IssueCertificate.bind(this);
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

function isObject(value: any): boolean {
  return typeof value === "object" && value !== null;
}

function isSet(value: any): boolean {
  return value !== null && value !== undefined;
}
