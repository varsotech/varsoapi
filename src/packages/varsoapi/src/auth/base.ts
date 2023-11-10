/* eslint-disable */
import * as _m0 from "protobufjs/minimal";

export const protobufPackage = "varso";

export interface User {
  uuid: string;
  discordUserId: string;
  banned: boolean;
}

function createBaseUser(): User {
  return { uuid: "", discordUserId: "", banned: false };
}

export const User = {
  encode(message: User, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.uuid !== "") {
      writer.uint32(10).string(message.uuid);
    }
    if (message.discordUserId !== "") {
      writer.uint32(18).string(message.discordUserId);
    }
    if (message.banned === true) {
      writer.uint32(24).bool(message.banned);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): User {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseUser();
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

          message.discordUserId = reader.string();
          continue;
        case 3:
          if (tag !== 24) {
            break;
          }

          message.banned = reader.bool();
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): User {
    return {
      uuid: isSet(object.uuid) ? String(object.uuid) : "",
      discordUserId: isSet(object.discordUserId) ? String(object.discordUserId) : "",
      banned: isSet(object.banned) ? Boolean(object.banned) : false,
    };
  },

  toJSON(message: User): unknown {
    const obj: any = {};
    if (message.uuid !== "") {
      obj.uuid = message.uuid;
    }
    if (message.discordUserId !== "") {
      obj.discordUserId = message.discordUserId;
    }
    if (message.banned === true) {
      obj.banned = message.banned;
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<User>, I>>(base?: I): User {
    return User.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<User>, I>>(object: I): User {
    const message = createBaseUser();
    message.uuid = object.uuid ?? "";
    message.discordUserId = object.discordUserId ?? "";
    message.banned = object.banned ?? false;
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
