/* eslint-disable */
import * as Long from "long";
import { util, configure, Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "archive.cda";

export interface CDA {
  creator: string;
  id: number;
  cid: string;
  ownership: { [key: string]: number };
  expiration: number;
}

export interface CDA_OwnershipEntry {
  key: string;
  value: number;
}

const baseCDA: object = { creator: "", id: 0, cid: "", expiration: 0 };

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
    Object.entries(message.ownership).forEach(([key, value]) => {
      CDA_OwnershipEntry.encode(
        { key: key as any, value },
        writer.uint32(34).fork()
      ).ldelim();
    });
    if (message.expiration !== 0) {
      writer.uint32(40).uint64(message.expiration);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): CDA {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseCDA } as CDA;
    message.ownership = {};
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
          const entry4 = CDA_OwnershipEntry.decode(reader, reader.uint32());
          if (entry4.value !== undefined) {
            message.ownership[entry4.key] = entry4.value;
          }
          break;
        case 5:
          message.expiration = longToNumber(reader.uint64() as Long);
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
    message.ownership = {};
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
      Object.entries(object.ownership).forEach(([key, value]) => {
        message.ownership[key] = Number(value);
      });
    }
    if (object.expiration !== undefined && object.expiration !== null) {
      message.expiration = Number(object.expiration);
    } else {
      message.expiration = 0;
    }
    return message;
  },

  toJSON(message: CDA): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.id !== undefined && (obj.id = message.id);
    message.cid !== undefined && (obj.cid = message.cid);
    obj.ownership = {};
    if (message.ownership) {
      Object.entries(message.ownership).forEach(([k, v]) => {
        obj.ownership[k] = v;
      });
    }
    message.expiration !== undefined && (obj.expiration = message.expiration);
    return obj;
  },

  fromPartial(object: DeepPartial<CDA>): CDA {
    const message = { ...baseCDA } as CDA;
    message.ownership = {};
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
      Object.entries(object.ownership).forEach(([key, value]) => {
        if (value !== undefined) {
          message.ownership[key] = Number(value);
        }
      });
    }
    if (object.expiration !== undefined && object.expiration !== null) {
      message.expiration = object.expiration;
    } else {
      message.expiration = 0;
    }
    return message;
  },
};

const baseCDA_OwnershipEntry: object = { key: "", value: 0 };

export const CDA_OwnershipEntry = {
  encode(
    message: CDA_OwnershipEntry,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.key !== "") {
      writer.uint32(10).string(message.key);
    }
    if (message.value !== 0) {
      writer.uint32(16).uint64(message.value);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): CDA_OwnershipEntry {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseCDA_OwnershipEntry } as CDA_OwnershipEntry;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.key = reader.string();
          break;
        case 2:
          message.value = longToNumber(reader.uint64() as Long);
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): CDA_OwnershipEntry {
    const message = { ...baseCDA_OwnershipEntry } as CDA_OwnershipEntry;
    if (object.key !== undefined && object.key !== null) {
      message.key = String(object.key);
    } else {
      message.key = "";
    }
    if (object.value !== undefined && object.value !== null) {
      message.value = Number(object.value);
    } else {
      message.value = 0;
    }
    return message;
  },

  toJSON(message: CDA_OwnershipEntry): unknown {
    const obj: any = {};
    message.key !== undefined && (obj.key = message.key);
    message.value !== undefined && (obj.value = message.value);
    return obj;
  },

  fromPartial(object: DeepPartial<CDA_OwnershipEntry>): CDA_OwnershipEntry {
    const message = { ...baseCDA_OwnershipEntry } as CDA_OwnershipEntry;
    if (object.key !== undefined && object.key !== null) {
      message.key = object.key;
    } else {
      message.key = "";
    }
    if (object.value !== undefined && object.value !== null) {
      message.value = object.value;
    } else {
      message.value = 0;
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
