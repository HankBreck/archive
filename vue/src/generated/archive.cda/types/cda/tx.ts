/* eslint-disable */
import { Reader, util, configure, Writer } from "protobufjs/minimal";
import * as Long from "long";
import { Ownership } from "../cda/cda";

export const protobufPackage = "archive.cda";

export interface MsgCreateCDA {
  creator: string;
  cid: string;
  ownership: Ownership[];
  expiration: number;
}

export interface MsgCreateCDAResponse {
  id: number;
}

export interface MsgApproveCda {
  creator: string;
  cdaId: number;
  ownership: Ownership[];
}

export interface MsgApproveCdaResponse {}

export interface MsgFinalizeCda {
  creator: string;
  cdaId: number;
}

export interface MsgFinalizeCdaResponse {}

const baseMsgCreateCDA: object = { creator: "", cid: "", expiration: 0 };

export const MsgCreateCDA = {
  encode(message: MsgCreateCDA, writer: Writer = Writer.create()): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.cid !== "") {
      writer.uint32(18).string(message.cid);
    }
    for (const v of message.ownership) {
      Ownership.encode(v!, writer.uint32(26).fork()).ldelim();
    }
    if (message.expiration !== 0) {
      writer.uint32(32).uint64(message.expiration);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgCreateCDA {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgCreateCDA } as MsgCreateCDA;
    message.ownership = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.cid = reader.string();
          break;
        case 3:
          message.ownership.push(Ownership.decode(reader, reader.uint32()));
          break;
        case 4:
          message.expiration = longToNumber(reader.uint64() as Long);
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgCreateCDA {
    const message = { ...baseMsgCreateCDA } as MsgCreateCDA;
    message.ownership = [];
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (object.cid !== undefined && object.cid !== null) {
      message.cid = String(object.cid);
    } else {
      message.cid = "";
    }
    if (object.ownership !== undefined && object.ownership !== null) {
      for (const e of object.ownership) {
        message.ownership.push(Ownership.fromJSON(e));
      }
    }
    if (object.expiration !== undefined && object.expiration !== null) {
      message.expiration = Number(object.expiration);
    } else {
      message.expiration = 0;
    }
    return message;
  },

  toJSON(message: MsgCreateCDA): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.cid !== undefined && (obj.cid = message.cid);
    if (message.ownership) {
      obj.ownership = message.ownership.map((e) =>
        e ? Ownership.toJSON(e) : undefined
      );
    } else {
      obj.ownership = [];
    }
    message.expiration !== undefined && (obj.expiration = message.expiration);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgCreateCDA>): MsgCreateCDA {
    const message = { ...baseMsgCreateCDA } as MsgCreateCDA;
    message.ownership = [];
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (object.cid !== undefined && object.cid !== null) {
      message.cid = object.cid;
    } else {
      message.cid = "";
    }
    if (object.ownership !== undefined && object.ownership !== null) {
      for (const e of object.ownership) {
        message.ownership.push(Ownership.fromPartial(e));
      }
    }
    if (object.expiration !== undefined && object.expiration !== null) {
      message.expiration = object.expiration;
    } else {
      message.expiration = 0;
    }
    return message;
  },
};

const baseMsgCreateCDAResponse: object = { id: 0 };

export const MsgCreateCDAResponse = {
  encode(
    message: MsgCreateCDAResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.id !== 0) {
      writer.uint32(8).uint64(message.id);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgCreateCDAResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgCreateCDAResponse } as MsgCreateCDAResponse;
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

  fromJSON(object: any): MsgCreateCDAResponse {
    const message = { ...baseMsgCreateCDAResponse } as MsgCreateCDAResponse;
    if (object.id !== undefined && object.id !== null) {
      message.id = Number(object.id);
    } else {
      message.id = 0;
    }
    return message;
  },

  toJSON(message: MsgCreateCDAResponse): unknown {
    const obj: any = {};
    message.id !== undefined && (obj.id = message.id);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgCreateCDAResponse>): MsgCreateCDAResponse {
    const message = { ...baseMsgCreateCDAResponse } as MsgCreateCDAResponse;
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
    for (const v of message.ownership) {
      Ownership.encode(v!, writer.uint32(26).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgApproveCda {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgApproveCda } as MsgApproveCda;
    message.ownership = [];
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
          message.ownership.push(Ownership.decode(reader, reader.uint32()));
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
    message.ownership = [];
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
    if (object.ownership !== undefined && object.ownership !== null) {
      for (const e of object.ownership) {
        message.ownership.push(Ownership.fromJSON(e));
      }
    }
    return message;
  },

  toJSON(message: MsgApproveCda): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.cdaId !== undefined && (obj.cdaId = message.cdaId);
    if (message.ownership) {
      obj.ownership = message.ownership.map((e) =>
        e ? Ownership.toJSON(e) : undefined
      );
    } else {
      obj.ownership = [];
    }
    return obj;
  },

  fromPartial(object: DeepPartial<MsgApproveCda>): MsgApproveCda {
    const message = { ...baseMsgApproveCda } as MsgApproveCda;
    message.ownership = [];
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
    if (object.ownership !== undefined && object.ownership !== null) {
      for (const e of object.ownership) {
        message.ownership.push(Ownership.fromPartial(e));
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
  CreateCDA(request: MsgCreateCDA): Promise<MsgCreateCDAResponse>;
  ApproveCda(request: MsgApproveCda): Promise<MsgApproveCdaResponse>;
  /** this line is used by starport scaffolding # proto/tx/rpc */
  FinalizeCda(request: MsgFinalizeCda): Promise<MsgFinalizeCdaResponse>;
}

export class MsgClientImpl implements Msg {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
  }
  CreateCDA(request: MsgCreateCDA): Promise<MsgCreateCDAResponse> {
    const data = MsgCreateCDA.encode(request).finish();
    const promise = this.rpc.request("archive.cda.Msg", "CreateCDA", data);
    return promise.then((data) =>
      MsgCreateCDAResponse.decode(new Reader(data))
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
