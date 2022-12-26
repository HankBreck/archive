/* eslint-disable */
import Long from "long";
import _m0 from "protobufjs/minimal";

export const protobufPackage = "archive.identity";

/** HashEntry is a record of a field: hash mapping used in Certificates */
export interface HashEntry {
  field: string;
  hash: string;
}

/** Certificate is the structure of an identity certificate */
export interface Certificate {
  id: number;
  issuerAddress: string;
  salt: string;
  metadataSchemaUri: string;
  hashes: HashEntry[];
}

function createBaseHashEntry(): HashEntry {
  return { field: "", hash: "" };
}

export const HashEntry = {
  encode(message: HashEntry, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.field !== "") {
      writer.uint32(10).string(message.field);
    }
    if (message.hash !== "") {
      writer.uint32(18).string(message.hash);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): HashEntry {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseHashEntry();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.field = reader.string();
          break;
        case 2:
          message.hash = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): HashEntry {
    return {
      field: isSet(object.field) ? String(object.field) : "",
      hash: isSet(object.hash) ? String(object.hash) : "",
    };
  },

  toJSON(message: HashEntry): unknown {
    const obj: any = {};
    message.field !== undefined && (obj.field = message.field);
    message.hash !== undefined && (obj.hash = message.hash);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<HashEntry>, I>>(object: I): HashEntry {
    const message = createBaseHashEntry();
    message.field = object.field ?? "";
    message.hash = object.hash ?? "";
    return message;
  },
};

function createBaseCertificate(): Certificate {
  return { id: 0, issuerAddress: "", salt: "", metadataSchemaUri: "", hashes: [] };
}

export const Certificate = {
  encode(message: Certificate, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.id !== 0) {
      writer.uint32(8).uint64(message.id);
    }
    if (message.issuerAddress !== "") {
      writer.uint32(18).string(message.issuerAddress);
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

  decode(input: _m0.Reader | Uint8Array, length?: number): Certificate {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseCertificate();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.id = longToNumber(reader.uint64() as Long);
          break;
        case 2:
          message.issuerAddress = reader.string();
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

  fromJSON(object: any): Certificate {
    return {
      id: isSet(object.id) ? Number(object.id) : 0,
      issuerAddress: isSet(object.issuerAddress) ? String(object.issuerAddress) : "",
      salt: isSet(object.salt) ? String(object.salt) : "",
      metadataSchemaUri: isSet(object.metadataSchemaUri) ? String(object.metadataSchemaUri) : "",
      hashes: Array.isArray(object?.hashes) ? object.hashes.map((e: any) => HashEntry.fromJSON(e)) : [],
    };
  },

  toJSON(message: Certificate): unknown {
    const obj: any = {};
    message.id !== undefined && (obj.id = Math.round(message.id));
    message.issuerAddress !== undefined && (obj.issuerAddress = message.issuerAddress);
    message.salt !== undefined && (obj.salt = message.salt);
    message.metadataSchemaUri !== undefined && (obj.metadataSchemaUri = message.metadataSchemaUri);
    if (message.hashes) {
      obj.hashes = message.hashes.map((e) => e ? HashEntry.toJSON(e) : undefined);
    } else {
      obj.hashes = [];
    }
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<Certificate>, I>>(object: I): Certificate {
    const message = createBaseCertificate();
    message.id = object.id ?? 0;
    message.issuerAddress = object.issuerAddress ?? "";
    message.salt = object.salt ?? "";
    message.metadataSchemaUri = object.metadataSchemaUri ?? "";
    message.hashes = object.hashes?.map((e) => HashEntry.fromPartial(e)) || [];
    return message;
  },
};

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
