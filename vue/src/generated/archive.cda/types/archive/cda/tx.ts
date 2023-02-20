/* eslint-disable */
import Long from "long";
import _m0 from "protobufjs/minimal";
import { Timestamp } from "../../google/protobuf/timestamp";

export const protobufPackage = "archive.cda";

export interface MsgCreateCda {
  creator: string;
  signingParties: string[];
  contractId: number;
  legalMetadataUri: string;
  signingData: Uint8Array;
  utcExpireTime: Date | undefined;
}

export interface MsgCreateCdaResponse {
  id: number;
}

export interface MsgApproveCda {
  creator: string;
  cdaId: number;
  signingData: Uint8Array;
}

export interface MsgApproveCdaResponse {
}

export interface MsgFinalizeCda {
  creator: string;
  cdaId: number;
}

export interface MsgFinalizeCdaResponse {
}

function createBaseMsgCreateCda(): MsgCreateCda {
  return {
    creator: "",
    signingParties: [],
    contractId: 0,
    legalMetadataUri: "",
    signingData: new Uint8Array(),
    utcExpireTime: undefined,
  };
}

export const MsgCreateCda = {
  encode(message: MsgCreateCda, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    for (const v of message.signingParties) {
      writer.uint32(18).string(v!);
    }
    if (message.contractId !== 0) {
      writer.uint32(24).uint64(message.contractId);
    }
    if (message.legalMetadataUri !== "") {
      writer.uint32(34).string(message.legalMetadataUri);
    }
    if (message.signingData.length !== 0) {
      writer.uint32(42).bytes(message.signingData);
    }
    if (message.utcExpireTime !== undefined) {
      Timestamp.encode(toTimestamp(message.utcExpireTime), writer.uint32(50).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgCreateCda {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgCreateCda();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.signingParties.push(reader.string());
          break;
        case 3:
          message.contractId = longToNumber(reader.uint64() as Long);
          break;
        case 4:
          message.legalMetadataUri = reader.string();
          break;
        case 5:
          message.signingData = reader.bytes();
          break;
        case 6:
          message.utcExpireTime = fromTimestamp(Timestamp.decode(reader, reader.uint32()));
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgCreateCda {
    return {
      creator: isSet(object.creator) ? String(object.creator) : "",
      signingParties: Array.isArray(object?.signingParties) ? object.signingParties.map((e: any) => String(e)) : [],
      contractId: isSet(object.contractId) ? Number(object.contractId) : 0,
      legalMetadataUri: isSet(object.legalMetadataUri) ? String(object.legalMetadataUri) : "",
      signingData: isSet(object.signingData) ? bytesFromBase64(object.signingData) : new Uint8Array(),
      utcExpireTime: isSet(object.utcExpireTime) ? fromJsonTimestamp(object.utcExpireTime) : undefined,
    };
  },

  toJSON(message: MsgCreateCda): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    if (message.signingParties) {
      obj.signingParties = message.signingParties.map((e) => e);
    } else {
      obj.signingParties = [];
    }
    message.contractId !== undefined && (obj.contractId = Math.round(message.contractId));
    message.legalMetadataUri !== undefined && (obj.legalMetadataUri = message.legalMetadataUri);
    message.signingData !== undefined
      && (obj.signingData = base64FromBytes(
        message.signingData !== undefined ? message.signingData : new Uint8Array(),
      ));
    message.utcExpireTime !== undefined && (obj.utcExpireTime = message.utcExpireTime.toISOString());
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgCreateCda>, I>>(object: I): MsgCreateCda {
    const message = createBaseMsgCreateCda();
    message.creator = object.creator ?? "";
    message.signingParties = object.signingParties?.map((e) => e) || [];
    message.contractId = object.contractId ?? 0;
    message.legalMetadataUri = object.legalMetadataUri ?? "";
    message.signingData = object.signingData ?? new Uint8Array();
    message.utcExpireTime = object.utcExpireTime ?? undefined;
    return message;
  },
};

function createBaseMsgCreateCdaResponse(): MsgCreateCdaResponse {
  return { id: 0 };
}

export const MsgCreateCdaResponse = {
  encode(message: MsgCreateCdaResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.id !== 0) {
      writer.uint32(8).uint64(message.id);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgCreateCdaResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgCreateCdaResponse();
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

  fromJSON(object: any): MsgCreateCdaResponse {
    return { id: isSet(object.id) ? Number(object.id) : 0 };
  },

  toJSON(message: MsgCreateCdaResponse): unknown {
    const obj: any = {};
    message.id !== undefined && (obj.id = Math.round(message.id));
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgCreateCdaResponse>, I>>(object: I): MsgCreateCdaResponse {
    const message = createBaseMsgCreateCdaResponse();
    message.id = object.id ?? 0;
    return message;
  },
};

function createBaseMsgApproveCda(): MsgApproveCda {
  return { creator: "", cdaId: 0, signingData: new Uint8Array() };
}

export const MsgApproveCda = {
  encode(message: MsgApproveCda, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.cdaId !== 0) {
      writer.uint32(16).uint64(message.cdaId);
    }
    if (message.signingData.length !== 0) {
      writer.uint32(26).bytes(message.signingData);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgApproveCda {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgApproveCda();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.cdaId = longToNumber(reader.uint64() as Long);
          break;
        case 3:
          message.signingData = reader.bytes();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgApproveCda {
    return {
      creator: isSet(object.creator) ? String(object.creator) : "",
      cdaId: isSet(object.cdaId) ? Number(object.cdaId) : 0,
      signingData: isSet(object.signingData) ? bytesFromBase64(object.signingData) : new Uint8Array(),
    };
  },

  toJSON(message: MsgApproveCda): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.cdaId !== undefined && (obj.cdaId = Math.round(message.cdaId));
    message.signingData !== undefined
      && (obj.signingData = base64FromBytes(
        message.signingData !== undefined ? message.signingData : new Uint8Array(),
      ));
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgApproveCda>, I>>(object: I): MsgApproveCda {
    const message = createBaseMsgApproveCda();
    message.creator = object.creator ?? "";
    message.cdaId = object.cdaId ?? 0;
    message.signingData = object.signingData ?? new Uint8Array();
    return message;
  },
};

function createBaseMsgApproveCdaResponse(): MsgApproveCdaResponse {
  return {};
}

export const MsgApproveCdaResponse = {
  encode(_: MsgApproveCdaResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgApproveCdaResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgApproveCdaResponse();
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

  fromJSON(_: any): MsgApproveCdaResponse {
    return {};
  },

  toJSON(_: MsgApproveCdaResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgApproveCdaResponse>, I>>(_: I): MsgApproveCdaResponse {
    const message = createBaseMsgApproveCdaResponse();
    return message;
  },
};

function createBaseMsgFinalizeCda(): MsgFinalizeCda {
  return { creator: "", cdaId: 0 };
}

export const MsgFinalizeCda = {
  encode(message: MsgFinalizeCda, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.cdaId !== 0) {
      writer.uint32(16).uint64(message.cdaId);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgFinalizeCda {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgFinalizeCda();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.cdaId = longToNumber(reader.uint64() as Long);
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgFinalizeCda {
    return {
      creator: isSet(object.creator) ? String(object.creator) : "",
      cdaId: isSet(object.cdaId) ? Number(object.cdaId) : 0,
    };
  },

  toJSON(message: MsgFinalizeCda): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.cdaId !== undefined && (obj.cdaId = Math.round(message.cdaId));
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgFinalizeCda>, I>>(object: I): MsgFinalizeCda {
    const message = createBaseMsgFinalizeCda();
    message.creator = object.creator ?? "";
    message.cdaId = object.cdaId ?? 0;
    return message;
  },
};

function createBaseMsgFinalizeCdaResponse(): MsgFinalizeCdaResponse {
  return {};
}

export const MsgFinalizeCdaResponse = {
  encode(_: MsgFinalizeCdaResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgFinalizeCdaResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgFinalizeCdaResponse();
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

  fromJSON(_: any): MsgFinalizeCdaResponse {
    return {};
  },

  toJSON(_: MsgFinalizeCdaResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgFinalizeCdaResponse>, I>>(_: I): MsgFinalizeCdaResponse {
    const message = createBaseMsgFinalizeCdaResponse();
    return message;
  },
};

/** Msg defines the Msg service. */
export interface Msg {
  CreateCda(request: MsgCreateCda): Promise<MsgCreateCdaResponse>;
  ApproveCda(request: MsgApproveCda): Promise<MsgApproveCdaResponse>;
  /** this line is used by starport scaffolding # proto/tx/rpc */
  FinalizeCda(request: MsgFinalizeCda): Promise<MsgFinalizeCdaResponse>;
}

export class MsgClientImpl implements Msg {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
    this.CreateCda = this.CreateCda.bind(this);
    this.ApproveCda = this.ApproveCda.bind(this);
    this.FinalizeCda = this.FinalizeCda.bind(this);
  }
  CreateCda(request: MsgCreateCda): Promise<MsgCreateCdaResponse> {
    const data = MsgCreateCda.encode(request).finish();
    const promise = this.rpc.request("archive.cda.Msg", "CreateCda", data);
    return promise.then((data) => MsgCreateCdaResponse.decode(new _m0.Reader(data)));
  }

  ApproveCda(request: MsgApproveCda): Promise<MsgApproveCdaResponse> {
    const data = MsgApproveCda.encode(request).finish();
    const promise = this.rpc.request("archive.cda.Msg", "ApproveCda", data);
    return promise.then((data) => MsgApproveCdaResponse.decode(new _m0.Reader(data)));
  }

  FinalizeCda(request: MsgFinalizeCda): Promise<MsgFinalizeCdaResponse> {
    const data = MsgFinalizeCda.encode(request).finish();
    const promise = this.rpc.request("archive.cda.Msg", "FinalizeCda", data);
    return promise.then((data) => MsgFinalizeCdaResponse.decode(new _m0.Reader(data)));
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

function bytesFromBase64(b64: string): Uint8Array {
  if (globalThis.Buffer) {
    return Uint8Array.from(globalThis.Buffer.from(b64, "base64"));
  } else {
    const bin = globalThis.atob(b64);
    const arr = new Uint8Array(bin.length);
    for (let i = 0; i < bin.length; ++i) {
      arr[i] = bin.charCodeAt(i);
    }
    return arr;
  }
}

function base64FromBytes(arr: Uint8Array): string {
  if (globalThis.Buffer) {
    return globalThis.Buffer.from(arr).toString("base64");
  } else {
    const bin: string[] = [];
    arr.forEach((byte) => {
      bin.push(String.fromCharCode(byte));
    });
    return globalThis.btoa(bin.join(""));
  }
}

type Builtin = Date | Function | Uint8Array | string | number | boolean | undefined;

export type DeepPartial<T> = T extends Builtin ? T
  : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>>
  : T extends {} ? { [K in keyof T]?: DeepPartial<T[K]> }
  : Partial<T>;

type KeysOfUnion<T> = T extends T ? keyof T : never;
export type Exact<P, I extends P> = P extends Builtin ? P
  : P & { [K in keyof P]: Exact<P[K], I[K]> } & { [K in Exclude<keyof I, KeysOfUnion<P>>]: never };

function toTimestamp(date: Date): Timestamp {
  const seconds = date.getTime() / 1_000;
  const nanos = (date.getTime() % 1_000) * 1_000_000;
  return { seconds, nanos };
}

function fromTimestamp(t: Timestamp): Date {
  let millis = t.seconds * 1_000;
  millis += t.nanos / 1_000_000;
  return new Date(millis);
}

function fromJsonTimestamp(o: any): Date {
  if (o instanceof Date) {
    return o;
  } else if (typeof o === "string") {
    return new Date(o);
  } else {
    return fromTimestamp(Timestamp.fromJSON(o));
  }
}

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
