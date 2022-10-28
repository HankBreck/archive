/* eslint-disable */
import { Reader, util, configure, Writer } from "protobufjs/minimal";
import { Timestamp } from "../google/protobuf/timestamp";
import * as Long from "long";
import { Any } from "../google/protobuf/any";

export const protobufPackage = "archive.cda";

export interface MsgCreateCda {
  creator: string;
  signingParties: string[];
  contractId: number;
  legalMetadataUri: string;
  signingData: Any[];
  utcExpireTime: Date | undefined;
}

export interface MsgCreateCdaResponse {
  id: number;
}

export interface MsgApproveCda {
  creator: string;
  cdaId: number;
  signingData: Any[];
}

export interface MsgApproveCdaResponse {}

export interface MsgFinalizeCda {
  creator: string;
  cdaId: number;
}

export interface MsgFinalizeCdaResponse {}

const baseMsgCreateCda: object = {
  creator: "",
  signingParties: "",
  contractId: 0,
  legalMetadataUri: "",
};

export const MsgCreateCda = {
  encode(message: MsgCreateCda, writer: Writer = Writer.create()): Writer {
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
    for (const v of message.signingData) {
      Any.encode(v!, writer.uint32(42).fork()).ldelim();
    }
    if (message.utcExpireTime !== undefined) {
      Timestamp.encode(
        toTimestamp(message.utcExpireTime),
        writer.uint32(50).fork()
      ).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgCreateCda {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgCreateCda } as MsgCreateCda;
    message.signingParties = [];
    message.signingData = [];
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
          message.signingData.push(Any.decode(reader, reader.uint32()));
          break;
        case 6:
          message.utcExpireTime = fromTimestamp(
            Timestamp.decode(reader, reader.uint32())
          );
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgCreateCda {
    const message = { ...baseMsgCreateCda } as MsgCreateCda;
    message.signingParties = [];
    message.signingData = [];
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (object.signingParties !== undefined && object.signingParties !== null) {
      for (const e of object.signingParties) {
        message.signingParties.push(String(e));
      }
    }
    if (object.contractId !== undefined && object.contractId !== null) {
      message.contractId = Number(object.contractId);
    } else {
      message.contractId = 0;
    }
    if (
      object.legalMetadataUri !== undefined &&
      object.legalMetadataUri !== null
    ) {
      message.legalMetadataUri = String(object.legalMetadataUri);
    } else {
      message.legalMetadataUri = "";
    }
    if (object.signingData !== undefined && object.signingData !== null) {
      for (const e of object.signingData) {
        message.signingData.push(Any.fromJSON(e));
      }
    }
    if (object.utcExpireTime !== undefined && object.utcExpireTime !== null) {
      message.utcExpireTime = fromJsonTimestamp(object.utcExpireTime);
    } else {
      message.utcExpireTime = undefined;
    }
    return message;
  },

  toJSON(message: MsgCreateCda): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    if (message.signingParties) {
      obj.signingParties = message.signingParties.map((e) => e);
    } else {
      obj.signingParties = [];
    }
    message.contractId !== undefined && (obj.contractId = message.contractId);
    message.legalMetadataUri !== undefined &&
      (obj.legalMetadataUri = message.legalMetadataUri);
    if (message.signingData) {
      obj.signingData = message.signingData.map((e) =>
        e ? Any.toJSON(e) : undefined
      );
    } else {
      obj.signingData = [];
    }
    message.utcExpireTime !== undefined &&
      (obj.utcExpireTime =
        message.utcExpireTime !== undefined
          ? message.utcExpireTime.toISOString()
          : null);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgCreateCda>): MsgCreateCda {
    const message = { ...baseMsgCreateCda } as MsgCreateCda;
    message.signingParties = [];
    message.signingData = [];
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (object.signingParties !== undefined && object.signingParties !== null) {
      for (const e of object.signingParties) {
        message.signingParties.push(e);
      }
    }
    if (object.contractId !== undefined && object.contractId !== null) {
      message.contractId = object.contractId;
    } else {
      message.contractId = 0;
    }
    if (
      object.legalMetadataUri !== undefined &&
      object.legalMetadataUri !== null
    ) {
      message.legalMetadataUri = object.legalMetadataUri;
    } else {
      message.legalMetadataUri = "";
    }
    if (object.signingData !== undefined && object.signingData !== null) {
      for (const e of object.signingData) {
        message.signingData.push(Any.fromPartial(e));
      }
    }
    if (object.utcExpireTime !== undefined && object.utcExpireTime !== null) {
      message.utcExpireTime = object.utcExpireTime;
    } else {
      message.utcExpireTime = undefined;
    }
    return message;
  },
};

const baseMsgCreateCdaResponse: object = { id: 0 };

export const MsgCreateCdaResponse = {
  encode(
    message: MsgCreateCdaResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.id !== 0) {
      writer.uint32(8).uint64(message.id);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgCreateCdaResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgCreateCdaResponse } as MsgCreateCdaResponse;
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
    const message = { ...baseMsgCreateCdaResponse } as MsgCreateCdaResponse;
    if (object.id !== undefined && object.id !== null) {
      message.id = Number(object.id);
    } else {
      message.id = 0;
    }
    return message;
  },

  toJSON(message: MsgCreateCdaResponse): unknown {
    const obj: any = {};
    message.id !== undefined && (obj.id = message.id);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgCreateCdaResponse>): MsgCreateCdaResponse {
    const message = { ...baseMsgCreateCdaResponse } as MsgCreateCdaResponse;
    if (object.id !== undefined && object.id !== null) {
      message.id = object.id;
    } else {
      message.id = 0;
    }
    return message;
  },
};

const baseMsgApproveCda: object = { creator: "", cdaId: 0 };

export const MsgApproveCda = {
  encode(message: MsgApproveCda, writer: Writer = Writer.create()): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.cdaId !== 0) {
      writer.uint32(16).uint64(message.cdaId);
    }
    for (const v of message.signingData) {
      Any.encode(v!, writer.uint32(26).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgApproveCda {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgApproveCda } as MsgApproveCda;
    message.signingData = [];
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
          message.signingData.push(Any.decode(reader, reader.uint32()));
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgApproveCda {
    const message = { ...baseMsgApproveCda } as MsgApproveCda;
    message.signingData = [];
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (object.cdaId !== undefined && object.cdaId !== null) {
      message.cdaId = Number(object.cdaId);
    } else {
      message.cdaId = 0;
    }
    if (object.signingData !== undefined && object.signingData !== null) {
      for (const e of object.signingData) {
        message.signingData.push(Any.fromJSON(e));
      }
    }
    return message;
  },

  toJSON(message: MsgApproveCda): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.cdaId !== undefined && (obj.cdaId = message.cdaId);
    if (message.signingData) {
      obj.signingData = message.signingData.map((e) =>
        e ? Any.toJSON(e) : undefined
      );
    } else {
      obj.signingData = [];
    }
    return obj;
  },

  fromPartial(object: DeepPartial<MsgApproveCda>): MsgApproveCda {
    const message = { ...baseMsgApproveCda } as MsgApproveCda;
    message.signingData = [];
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (object.cdaId !== undefined && object.cdaId !== null) {
      message.cdaId = object.cdaId;
    } else {
      message.cdaId = 0;
    }
    if (object.signingData !== undefined && object.signingData !== null) {
      for (const e of object.signingData) {
        message.signingData.push(Any.fromPartial(e));
      }
    }
    return message;
  },
};

const baseMsgApproveCdaResponse: object = {};

export const MsgApproveCdaResponse = {
  encode(_: MsgApproveCdaResponse, writer: Writer = Writer.create()): Writer {
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgApproveCdaResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgApproveCdaResponse } as MsgApproveCdaResponse;
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
    const message = { ...baseMsgApproveCdaResponse } as MsgApproveCdaResponse;
    return message;
  },

  toJSON(_: MsgApproveCdaResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(_: DeepPartial<MsgApproveCdaResponse>): MsgApproveCdaResponse {
    const message = { ...baseMsgApproveCdaResponse } as MsgApproveCdaResponse;
    return message;
  },
};

const baseMsgFinalizeCda: object = { creator: "", cdaId: 0 };

export const MsgFinalizeCda = {
  encode(message: MsgFinalizeCda, writer: Writer = Writer.create()): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.cdaId !== 0) {
      writer.uint32(16).uint64(message.cdaId);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgFinalizeCda {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgFinalizeCda } as MsgFinalizeCda;
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
    const message = { ...baseMsgFinalizeCda } as MsgFinalizeCda;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (object.cdaId !== undefined && object.cdaId !== null) {
      message.cdaId = Number(object.cdaId);
    } else {
      message.cdaId = 0;
    }
    return message;
  },

  toJSON(message: MsgFinalizeCda): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.cdaId !== undefined && (obj.cdaId = message.cdaId);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgFinalizeCda>): MsgFinalizeCda {
    const message = { ...baseMsgFinalizeCda } as MsgFinalizeCda;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (object.cdaId !== undefined && object.cdaId !== null) {
      message.cdaId = object.cdaId;
    } else {
      message.cdaId = 0;
    }
    return message;
  },
};

const baseMsgFinalizeCdaResponse: object = {};

export const MsgFinalizeCdaResponse = {
  encode(_: MsgFinalizeCdaResponse, writer: Writer = Writer.create()): Writer {
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgFinalizeCdaResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgFinalizeCdaResponse } as MsgFinalizeCdaResponse;
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
    const message = { ...baseMsgFinalizeCdaResponse } as MsgFinalizeCdaResponse;
    return message;
  },

  toJSON(_: MsgFinalizeCdaResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(_: DeepPartial<MsgFinalizeCdaResponse>): MsgFinalizeCdaResponse {
    const message = { ...baseMsgFinalizeCdaResponse } as MsgFinalizeCdaResponse;
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
  }
  CreateCda(request: MsgCreateCda): Promise<MsgCreateCdaResponse> {
    const data = MsgCreateCda.encode(request).finish();
    const promise = this.rpc.request("archive.cda.Msg", "CreateCda", data);
    return promise.then((data) =>
      MsgCreateCdaResponse.decode(new Reader(data))
    );
  }

  ApproveCda(request: MsgApproveCda): Promise<MsgApproveCdaResponse> {
    const data = MsgApproveCda.encode(request).finish();
    const promise = this.rpc.request("archive.cda.Msg", "ApproveCda", data);
    return promise.then((data) =>
      MsgApproveCdaResponse.decode(new Reader(data))
    );
  }

  FinalizeCda(request: MsgFinalizeCda): Promise<MsgFinalizeCdaResponse> {
    const data = MsgFinalizeCda.encode(request).finish();
    const promise = this.rpc.request("archive.cda.Msg", "FinalizeCda", data);
    return promise.then((data) =>
      MsgFinalizeCdaResponse.decode(new Reader(data))
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

if (util.Long !== Long) {
  util.Long = Long as any;
  configure();
}
