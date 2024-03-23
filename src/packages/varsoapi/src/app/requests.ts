/* eslint-disable */
import _m0 from "protobufjs/minimal";
import { Organization, RSSFeed } from "./base";

export const protobufPackage = "varso";

export interface GetOrganizationsResponse {
  organizations: Organization[];
}

export interface GetNewsResponse {
  items: GetNewsResponseItem[];
}

export interface GetNewsResponseItem {
  feed: RSSFeed | undefined;
  organization: Organization | undefined;
}

function createBaseGetOrganizationsResponse(): GetOrganizationsResponse {
  return { organizations: [] };
}

export const GetOrganizationsResponse = {
  encode(message: GetOrganizationsResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    for (const v of message.organizations) {
      Organization.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): GetOrganizationsResponse {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseGetOrganizationsResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.organizations.push(Organization.decode(reader, reader.uint32()));
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): GetOrganizationsResponse {
    return {
      organizations: Array.isArray(object?.organizations)
        ? object.organizations.map((e: any) => Organization.fromJSON(e))
        : [],
    };
  },

  toJSON(message: GetOrganizationsResponse): unknown {
    const obj: any = {};
    if (message.organizations?.length) {
      obj.organizations = message.organizations.map((e) => Organization.toJSON(e));
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<GetOrganizationsResponse>, I>>(base?: I): GetOrganizationsResponse {
    return GetOrganizationsResponse.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<GetOrganizationsResponse>, I>>(object: I): GetOrganizationsResponse {
    const message = createBaseGetOrganizationsResponse();
    message.organizations = object.organizations?.map((e) => Organization.fromPartial(e)) || [];
    return message;
  },
};

function createBaseGetNewsResponse(): GetNewsResponse {
  return { items: [] };
}

export const GetNewsResponse = {
  encode(message: GetNewsResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    for (const v of message.items) {
      GetNewsResponseItem.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): GetNewsResponse {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseGetNewsResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.items.push(GetNewsResponseItem.decode(reader, reader.uint32()));
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): GetNewsResponse {
    return { items: Array.isArray(object?.items) ? object.items.map((e: any) => GetNewsResponseItem.fromJSON(e)) : [] };
  },

  toJSON(message: GetNewsResponse): unknown {
    const obj: any = {};
    if (message.items?.length) {
      obj.items = message.items.map((e) => GetNewsResponseItem.toJSON(e));
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<GetNewsResponse>, I>>(base?: I): GetNewsResponse {
    return GetNewsResponse.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<GetNewsResponse>, I>>(object: I): GetNewsResponse {
    const message = createBaseGetNewsResponse();
    message.items = object.items?.map((e) => GetNewsResponseItem.fromPartial(e)) || [];
    return message;
  },
};

function createBaseGetNewsResponseItem(): GetNewsResponseItem {
  return { feed: undefined, organization: undefined };
}

export const GetNewsResponseItem = {
  encode(message: GetNewsResponseItem, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.feed !== undefined) {
      RSSFeed.encode(message.feed, writer.uint32(10).fork()).ldelim();
    }
    if (message.organization !== undefined) {
      Organization.encode(message.organization, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): GetNewsResponseItem {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseGetNewsResponseItem();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.feed = RSSFeed.decode(reader, reader.uint32());
          continue;
        case 2:
          if (tag !== 18) {
            break;
          }

          message.organization = Organization.decode(reader, reader.uint32());
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): GetNewsResponseItem {
    return {
      feed: isSet(object.feed) ? RSSFeed.fromJSON(object.feed) : undefined,
      organization: isSet(object.organization) ? Organization.fromJSON(object.organization) : undefined,
    };
  },

  toJSON(message: GetNewsResponseItem): unknown {
    const obj: any = {};
    if (message.feed !== undefined) {
      obj.feed = RSSFeed.toJSON(message.feed);
    }
    if (message.organization !== undefined) {
      obj.organization = Organization.toJSON(message.organization);
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<GetNewsResponseItem>, I>>(base?: I): GetNewsResponseItem {
    return GetNewsResponseItem.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<GetNewsResponseItem>, I>>(object: I): GetNewsResponseItem {
    const message = createBaseGetNewsResponseItem();
    message.feed = (object.feed !== undefined && object.feed !== null) ? RSSFeed.fromPartial(object.feed) : undefined;
    message.organization = (object.organization !== undefined && object.organization !== null)
      ? Organization.fromPartial(object.organization)
      : undefined;
    return message;
  },
};

type Builtin = Date | Function | Uint8Array | string | number | boolean | undefined;

export type DeepPartial<T> = T extends Builtin ? T
  : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>>
  : T extends {} ? { [K in keyof T]?: DeepPartial<T[K]> }
  : Partial<T>;

type KeysOfUnion<T> = T extends T ? keyof T : never;
export type Exact<P, I extends P> = P extends Builtin ? P
  : P & { [K in keyof P]: Exact<P[K], I[K]> } & { [K in Exclude<keyof I, KeysOfUnion<P>>]: never };

function isSet(value: any): boolean {
  return value !== null && value !== undefined;
}
