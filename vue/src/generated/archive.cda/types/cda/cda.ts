/* eslint-disable */
import * as Long from "long";
import { util, configure, Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "archive.cda";

export interface CDA {
  creator: string;
  id: number;
  cid: string;
  ownership: Ownership[];
  expiration: number;
  approved: boolean;
}

export interface Ownership {
  owner: string;
  ownership: number;
}

const baseCDA: object = {
  creator: "",
  id: 0,
  cid: "",
  expiration: 0,
  approved: false,
};

export const CDA = {
  encode(message: CDA, writer: Writer = Writer.create()): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.id !== 0) {
      writer.uint32(16).uint64(message.id);
    }
    if (message.cid !== "") {
      writer.uint32(26).string(message.cid);
    }
    for (const v of message.ownership) {
      Ownership.encode(v!, writer.uint32(34).fork()).ldelim();
    }
    if (message.expiration !== 0) {
      writer.uint32(40).uint64(message.expiration);
    }
    if (message.approved === true) {
      writer.uint32(48).bool(message.approved);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): CDA {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseCDA } as CDA;
    message.ownership = [];
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
          message.cid = reader.string();
          break;
        case 4:
          message.ownership.push(Ownership.decode(reader, reader.uint32()));
          break;
        case 5:
          message.expiration = longToNumber(reader.uint64() as Long);
          break;
        case 6:
          message.approved = reader.bool();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): CDA {
    const message = { ...baseCDA } as CDA;
    message.ownership = [];
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (object.id !== undefined && object.id !== null) {
      message.id = Number(object.id);
    } else {
      message.id = 0;
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
    if (object.approved !== undefined && object.approved !== null) {
      message.approved = Boolean(object.approved);
    } else {
      message.approved = false;
    }
    return message;
  },

  toJSON(message: CDA): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.id !== undefined && (obj.id = message.id);
    message.cid !== undefined && (obj.cid = message.cid);
    if (message.ownership) {
      obj.ownership = message.ownership.map((e) =>
        e ? Ownership.toJSON(e) : undefined
      );
    } else {
      obj.ownership = [];
    }
    message.expiration !== undefined && (obj.expiration = message.expiration);
    message.approved !== undefined && (obj.approved = message.approved);
    return obj;
  },

  fromPartial(object: DeepPartial<CDA>): CDA {
    const message = { ...baseCDA } as CDA;
    message.ownership = [];
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (object.id !== undefined && object.id !== null) {
      message.id = object.id;
    } else {
      message.id = 0;
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
    if (object.approved !== undefined && object.approved !== null) {
      message.approved = object.approved;
    } else {
      message.approved = false;
    }
    return message;
  },
};

const baseOwnership: object = { owner: "", ownership: 0 };

export const Ownership = {
  encode(message: Ownership, writer: Writer = Writer.create()): Writer {
    if (message.owner !== "") {
      writer.uint32(10).string(message.owner);
    }
    if (message.ownership !== 0) {
      writer.uint32(16).uint64(message.ownership);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): Ownership {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseOwnership } as Ownership;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.owner = reader.string();
          break;
        case 2:
          message.ownership = longToNumber(reader.uint64() as Long);
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): Ownership {
    const message = { ...baseOwnership } as Ownership;
    if (object.owner !== undefined && object.owner !== null) {
      message.owner = String(object.owner);
    } else {
      message.owner = "";
    }
    if (object.ownership !== undefined && object.ownership !== null) {
      message.ownership = Number(object.ownership);
    } else {
      message.ownership = 0;
    }
    return message;
  },

  toJSON(message: Ownership): unknown {
    const obj: any = {};
    message.owner !== undefined && (obj.owner = message.owner);
    message.ownership !== undefined && (obj.ownership = message.ownership);
    return obj;
  },

  fromPartial(object: DeepPartial<Ownership>): Ownership {
    const message = { ...baseOwnership } as Ownership;
    if (object.owner !== undefined && object.owner !== null) {
      message.owner = object.owner;
    } else {
      message.owner = "";
    }
    if (object.ownership !== undefined && object.ownership !== null) {
      message.ownership = object.ownership;
    } else {
      message.ownership = 0;
    }
    return message;
  },
};

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
