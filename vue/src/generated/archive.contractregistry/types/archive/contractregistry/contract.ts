/* eslint-disable */
import Long from "long";
import _m0 from "protobufjs/minimal";

export const protobufPackage = "archive.contractregistry";

export enum ContactMethod {
  /** Phone - People won't want to publish their phone number on here */
  Phone = 0,
  Email = 1,
  UNRECOGNIZED = -1,
}

export function contactMethodFromJSON(object: any): ContactMethod {
  switch (object) {
    case 0:
    case "Phone":
      return ContactMethod.Phone;
    case 1:
    case "Email":
      return ContactMethod.Email;
    case -1:
    case "UNRECOGNIZED":
    default:
      return ContactMethod.UNRECOGNIZED;
  }
}

export function contactMethodToJSON(object: ContactMethod): string {
  switch (object) {
    case ContactMethod.Phone:
      return "Phone";
    case ContactMethod.Email:
      return "Email";
    case ContactMethod.UNRECOGNIZED:
    default:
      return "UNRECOGNIZED";
  }
}

export interface Contract {
  id: number;
  description: string;
  authors: string[];
  contactInfo: ContactInfo | undefined;
  moreInfoUri: string;
  templateUri: string;
  templateSchemaUri: string;
}

export interface ContactInfo {
  method: ContactMethod;
  value: string;
}

function createBaseContract(): Contract {
  return {
    id: 0,
    description: "",
    authors: [],
    contactInfo: undefined,
    moreInfoUri: "",
    templateUri: "",
    templateSchemaUri: "",
  };
}

export const Contract = {
  encode(message: Contract, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.id !== 0) {
      writer.uint32(8).uint64(message.id);
    }
    if (message.description !== "") {
      writer.uint32(18).string(message.description);
    }
    for (const v of message.authors) {
      writer.uint32(26).string(v!);
    }
    if (message.contactInfo !== undefined) {
      ContactInfo.encode(message.contactInfo, writer.uint32(34).fork()).ldelim();
    }
    if (message.moreInfoUri !== "") {
      writer.uint32(42).string(message.moreInfoUri);
    }
    if (message.templateUri !== "") {
      writer.uint32(50).string(message.templateUri);
    }
    if (message.templateSchemaUri !== "") {
      writer.uint32(58).string(message.templateSchemaUri);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): Contract {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseContract();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.id = longToNumber(reader.uint64() as Long);
          break;
        case 2:
          message.description = reader.string();
          break;
        case 3:
          message.authors.push(reader.string());
          break;
        case 4:
          message.contactInfo = ContactInfo.decode(reader, reader.uint32());
          break;
        case 5:
          message.moreInfoUri = reader.string();
          break;
        case 6:
          message.templateUri = reader.string();
          break;
        case 7:
          message.templateSchemaUri = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): Contract {
    return {
      id: isSet(object.id) ? Number(object.id) : 0,
      description: isSet(object.description) ? String(object.description) : "",
      authors: Array.isArray(object?.authors) ? object.authors.map((e: any) => String(e)) : [],
      contactInfo: isSet(object.contactInfo) ? ContactInfo.fromJSON(object.contactInfo) : undefined,
      moreInfoUri: isSet(object.moreInfoUri) ? String(object.moreInfoUri) : "",
      templateUri: isSet(object.templateUri) ? String(object.templateUri) : "",
      templateSchemaUri: isSet(object.templateSchemaUri) ? String(object.templateSchemaUri) : "",
    };
  },

  toJSON(message: Contract): unknown {
    const obj: any = {};
    message.id !== undefined && (obj.id = Math.round(message.id));
    message.description !== undefined && (obj.description = message.description);
    if (message.authors) {
      obj.authors = message.authors.map((e) => e);
    } else {
      obj.authors = [];
    }
    message.contactInfo !== undefined
      && (obj.contactInfo = message.contactInfo ? ContactInfo.toJSON(message.contactInfo) : undefined);
    message.moreInfoUri !== undefined && (obj.moreInfoUri = message.moreInfoUri);
    message.templateUri !== undefined && (obj.templateUri = message.templateUri);
    message.templateSchemaUri !== undefined && (obj.templateSchemaUri = message.templateSchemaUri);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<Contract>, I>>(object: I): Contract {
    const message = createBaseContract();
    message.id = object.id ?? 0;
    message.description = object.description ?? "";
    message.authors = object.authors?.map((e) => e) || [];
    message.contactInfo = (object.contactInfo !== undefined && object.contactInfo !== null)
      ? ContactInfo.fromPartial(object.contactInfo)
      : undefined;
    message.moreInfoUri = object.moreInfoUri ?? "";
    message.templateUri = object.templateUri ?? "";
    message.templateSchemaUri = object.templateSchemaUri ?? "";
    return message;
  },
};

function createBaseContactInfo(): ContactInfo {
  return { method: 0, value: "" };
}

export const ContactInfo = {
  encode(message: ContactInfo, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.method !== 0) {
      writer.uint32(8).int32(message.method);
    }
    if (message.value !== "") {
      writer.uint32(18).string(message.value);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): ContactInfo {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseContactInfo();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.method = reader.int32() as any;
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

  fromJSON(object: any): ContactInfo {
    return {
      method: isSet(object.method) ? contactMethodFromJSON(object.method) : 0,
      value: isSet(object.value) ? String(object.value) : "",
    };
  },

  toJSON(message: ContactInfo): unknown {
    const obj: any = {};
    message.method !== undefined && (obj.method = contactMethodToJSON(message.method));
    message.value !== undefined && (obj.value = message.value);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<ContactInfo>, I>>(object: I): ContactInfo {
    const message = createBaseContactInfo();
    message.method = object.method ?? 0;
    message.value = object.value ?? "";
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
