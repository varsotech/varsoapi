/* eslint-disable */
import Long from "long";
import _m0 from "protobufjs/minimal";

export const protobufPackage = "varso";

export interface Post {
  uuid: string;
  title: string;
  coverImageUrl: string;
  totalVotes: number;
  commentsAmount: number;
}

export interface Album {
  name: string;
  artists: Artist[];
}

export interface Artist {
  name: string;
  spotifyLink: string;
}

function createBasePost(): Post {
  return { uuid: "", title: "", coverImageUrl: "", totalVotes: 0, commentsAmount: 0 };
}

export const Post = {
  encode(message: Post, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.uuid !== "") {
      writer.uint32(10).string(message.uuid);
    }
    if (message.title !== "") {
      writer.uint32(18).string(message.title);
    }
    if (message.coverImageUrl !== "") {
      writer.uint32(26).string(message.coverImageUrl);
    }
    if (message.totalVotes !== 0) {
      writer.uint32(32).int64(message.totalVotes);
    }
    if (message.commentsAmount !== 0) {
      writer.uint32(48).int64(message.commentsAmount);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): Post {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBasePost();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.uuid = reader.string();
          continue;
        case 2:
          if (tag !== 18) {
            break;
          }

          message.title = reader.string();
          continue;
        case 3:
          if (tag !== 26) {
            break;
          }

          message.coverImageUrl = reader.string();
          continue;
        case 4:
          if (tag !== 32) {
            break;
          }

          message.totalVotes = longToNumber(reader.int64() as Long);
          continue;
        case 6:
          if (tag !== 48) {
            break;
          }

          message.commentsAmount = longToNumber(reader.int64() as Long);
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): Post {
    return {
      uuid: isSet(object.uuid) ? String(object.uuid) : "",
      title: isSet(object.title) ? String(object.title) : "",
      coverImageUrl: isSet(object.coverImageUrl) ? String(object.coverImageUrl) : "",
      totalVotes: isSet(object.totalVotes) ? Number(object.totalVotes) : 0,
      commentsAmount: isSet(object.commentsAmount) ? Number(object.commentsAmount) : 0,
    };
  },

  toJSON(message: Post): unknown {
    const obj: any = {};
    if (message.uuid !== "") {
      obj.uuid = message.uuid;
    }
    if (message.title !== "") {
      obj.title = message.title;
    }
    if (message.coverImageUrl !== "") {
      obj.coverImageUrl = message.coverImageUrl;
    }
    if (message.totalVotes !== 0) {
      obj.totalVotes = Math.round(message.totalVotes);
    }
    if (message.commentsAmount !== 0) {
      obj.commentsAmount = Math.round(message.commentsAmount);
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<Post>, I>>(base?: I): Post {
    return Post.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<Post>, I>>(object: I): Post {
    const message = createBasePost();
    message.uuid = object.uuid ?? "";
    message.title = object.title ?? "";
    message.coverImageUrl = object.coverImageUrl ?? "";
    message.totalVotes = object.totalVotes ?? 0;
    message.commentsAmount = object.commentsAmount ?? 0;
    return message;
  },
};

function createBaseAlbum(): Album {
  return { name: "", artists: [] };
}

export const Album = {
  encode(message: Album, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.name !== "") {
      writer.uint32(10).string(message.name);
    }
    for (const v of message.artists) {
      Artist.encode(v!, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): Album {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseAlbum();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.name = reader.string();
          continue;
        case 2:
          if (tag !== 18) {
            break;
          }

          message.artists.push(Artist.decode(reader, reader.uint32()));
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): Album {
    return {
      name: isSet(object.name) ? String(object.name) : "",
      artists: Array.isArray(object?.artists) ? object.artists.map((e: any) => Artist.fromJSON(e)) : [],
    };
  },

  toJSON(message: Album): unknown {
    const obj: any = {};
    if (message.name !== "") {
      obj.name = message.name;
    }
    if (message.artists?.length) {
      obj.artists = message.artists.map((e) => Artist.toJSON(e));
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<Album>, I>>(base?: I): Album {
    return Album.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<Album>, I>>(object: I): Album {
    const message = createBaseAlbum();
    message.name = object.name ?? "";
    message.artists = object.artists?.map((e) => Artist.fromPartial(e)) || [];
    return message;
  },
};

function createBaseArtist(): Artist {
  return { name: "", spotifyLink: "" };
}

export const Artist = {
  encode(message: Artist, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.name !== "") {
      writer.uint32(10).string(message.name);
    }
    if (message.spotifyLink !== "") {
      writer.uint32(18).string(message.spotifyLink);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): Artist {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseArtist();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.name = reader.string();
          continue;
        case 2:
          if (tag !== 18) {
            break;
          }

          message.spotifyLink = reader.string();
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): Artist {
    return {
      name: isSet(object.name) ? String(object.name) : "",
      spotifyLink: isSet(object.spotifyLink) ? String(object.spotifyLink) : "",
    };
  },

  toJSON(message: Artist): unknown {
    const obj: any = {};
    if (message.name !== "") {
      obj.name = message.name;
    }
    if (message.spotifyLink !== "") {
      obj.spotifyLink = message.spotifyLink;
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<Artist>, I>>(base?: I): Artist {
    return Artist.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<Artist>, I>>(object: I): Artist {
    const message = createBaseArtist();
    message.name = object.name ?? "";
    message.spotifyLink = object.spotifyLink ?? "";
    return message;
  },
};

declare const self: any | undefined;
declare const window: any | undefined;
declare const global: any | undefined;
const tsProtoGlobalThis: any = (() => {
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
    throw new tsProtoGlobalThis.Error("Value is larger than Number.MAX_SAFE_INTEGER");
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
