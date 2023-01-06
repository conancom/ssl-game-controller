/* eslint-disable */
import { Timestamp } from "./google/protobuf/timestamp";
import { RobotId, Team, teamFromJSON, teamToJSON } from "./ssl_gc_common";
import { Vector2, Vector3 } from "./ssl_gc_geometry";

/** The GC state contains settings and state independent of the match state */
export interface GcState {
  /** the state of each team */
  teamState: { [key: string]: GcStateTeam };
  /** the states of the auto referees */
  autoRefState: { [key: string]: GcStateAutoRef };
  /** the attached trackers (uuid -> source_name) */
  trackers: { [key: string]: string };
  /** the next actions that can be executed when continuing */
  continueActions: ContinueAction[];
}

export interface GcState_TeamStateEntry {
  key: string;
  value: GcStateTeam | undefined;
}

export interface GcState_AutoRefStateEntry {
  key: string;
  value: GcStateAutoRef | undefined;
}

export interface GcState_TrackersEntry {
  key: string;
  value: string;
}

/** The GC state for a single team */
export interface GcStateTeam {
  /** true: The team is connected */
  connected: boolean;
  /** true: The team connected via TLS with a verified certificate */
  connectionVerified: boolean;
  /** true: The remote control for the team is connected */
  remoteControlConnected: boolean;
  /** true: The remote control for the team connected via TLS with a verified certificate */
  remoteControlConnectionVerified: boolean;
  /** the advantage choice of the team */
  advantageChoice: TeamAdvantageChoice | undefined;
}

/** The choice from a team regarding the advantage rule */
export interface TeamAdvantageChoice {
  /** the choice of the team */
  choice: TeamAdvantageChoice_AdvantageChoice;
}

/** possible advantage choices */
export enum TeamAdvantageChoice_AdvantageChoice {
  /** STOP - stop the game */
  STOP = "STOP",
  /** CONTINUE - keep the match running */
  CONTINUE = "CONTINUE",
  UNRECOGNIZED = "UNRECOGNIZED",
}

export function teamAdvantageChoice_AdvantageChoiceFromJSON(object: any): TeamAdvantageChoice_AdvantageChoice {
  switch (object) {
    case 0:
    case "STOP":
      return TeamAdvantageChoice_AdvantageChoice.STOP;
    case 1:
    case "CONTINUE":
      return TeamAdvantageChoice_AdvantageChoice.CONTINUE;
    case -1:
    case "UNRECOGNIZED":
    default:
      return TeamAdvantageChoice_AdvantageChoice.UNRECOGNIZED;
  }
}

export function teamAdvantageChoice_AdvantageChoiceToJSON(object: TeamAdvantageChoice_AdvantageChoice): string {
  switch (object) {
    case TeamAdvantageChoice_AdvantageChoice.STOP:
      return "STOP";
    case TeamAdvantageChoice_AdvantageChoice.CONTINUE:
      return "CONTINUE";
    case TeamAdvantageChoice_AdvantageChoice.UNRECOGNIZED:
    default:
      return "UNRECOGNIZED";
  }
}

/** The GC state of an auto referee */
export interface GcStateAutoRef {
  /** true: The autoRef connected via TLS with a verified certificate */
  connectionVerified: boolean;
}

/** GC state of a tracker */
export interface GcStateTracker {
  /** Name of the source */
  sourceName: string;
  /** UUID of the source */
  uuid: string;
  /** Current ball */
  ball:
    | Ball
    | undefined;
  /** Current robots */
  robots: Robot[];
}

/** The ball state */
export interface Ball {
  /** ball position [m] */
  pos:
    | Vector3
    | undefined;
  /** ball velocity [m/s] */
  vel: Vector3 | undefined;
}

/** The robot state */
export interface Robot {
  /** robot id and team */
  id:
    | RobotId
    | undefined;
  /** robot position [m] */
  pos: Vector2 | undefined;
}

export interface ContinueAction {
  /** type of action that will be performed next */
  type: ContinueAction_Type;
  /** for which team (if team specific) */
  forTeam: Team;
  /** list of issues that hinders the game from continuing */
  continuationIssues: string[];
  /** timestamp at which the action will be ready (to give some preparation time) */
  readyAt:
    | Date
    | undefined;
  /** state of the action */
  state: ContinueAction_State;
}

export enum ContinueAction_Type {
  TYPE_UNKNOWN = "TYPE_UNKNOWN",
  HALT = "HALT",
  RESUME_FROM_HALT = "RESUME_FROM_HALT",
  STOP_GAME = "STOP_GAME",
  RESUME_FROM_STOP = "RESUME_FROM_STOP",
  NEXT_COMMAND = "NEXT_COMMAND",
  BALL_PLACEMENT_START = "BALL_PLACEMENT_START",
  BALL_PLACEMENT_CANCEL = "BALL_PLACEMENT_CANCEL",
  TIMEOUT_START = "TIMEOUT_START",
  TIMEOUT_STOP = "TIMEOUT_STOP",
  BOT_SUBSTITUTION = "BOT_SUBSTITUTION",
  NEXT_STAGE = "NEXT_STAGE",
  UNRECOGNIZED = "UNRECOGNIZED",
}

export function continueAction_TypeFromJSON(object: any): ContinueAction_Type {
  switch (object) {
    case 0:
    case "TYPE_UNKNOWN":
      return ContinueAction_Type.TYPE_UNKNOWN;
    case 1:
    case "HALT":
      return ContinueAction_Type.HALT;
    case 10:
    case "RESUME_FROM_HALT":
      return ContinueAction_Type.RESUME_FROM_HALT;
    case 2:
    case "STOP_GAME":
      return ContinueAction_Type.STOP_GAME;
    case 11:
    case "RESUME_FROM_STOP":
      return ContinueAction_Type.RESUME_FROM_STOP;
    case 3:
    case "NEXT_COMMAND":
      return ContinueAction_Type.NEXT_COMMAND;
    case 4:
    case "BALL_PLACEMENT_START":
      return ContinueAction_Type.BALL_PLACEMENT_START;
    case 9:
    case "BALL_PLACEMENT_CANCEL":
      return ContinueAction_Type.BALL_PLACEMENT_CANCEL;
    case 5:
    case "TIMEOUT_START":
      return ContinueAction_Type.TIMEOUT_START;
    case 6:
    case "TIMEOUT_STOP":
      return ContinueAction_Type.TIMEOUT_STOP;
    case 7:
    case "BOT_SUBSTITUTION":
      return ContinueAction_Type.BOT_SUBSTITUTION;
    case 8:
    case "NEXT_STAGE":
      return ContinueAction_Type.NEXT_STAGE;
    case -1:
    case "UNRECOGNIZED":
    default:
      return ContinueAction_Type.UNRECOGNIZED;
  }
}

export function continueAction_TypeToJSON(object: ContinueAction_Type): string {
  switch (object) {
    case ContinueAction_Type.TYPE_UNKNOWN:
      return "TYPE_UNKNOWN";
    case ContinueAction_Type.HALT:
      return "HALT";
    case ContinueAction_Type.RESUME_FROM_HALT:
      return "RESUME_FROM_HALT";
    case ContinueAction_Type.STOP_GAME:
      return "STOP_GAME";
    case ContinueAction_Type.RESUME_FROM_STOP:
      return "RESUME_FROM_STOP";
    case ContinueAction_Type.NEXT_COMMAND:
      return "NEXT_COMMAND";
    case ContinueAction_Type.BALL_PLACEMENT_START:
      return "BALL_PLACEMENT_START";
    case ContinueAction_Type.BALL_PLACEMENT_CANCEL:
      return "BALL_PLACEMENT_CANCEL";
    case ContinueAction_Type.TIMEOUT_START:
      return "TIMEOUT_START";
    case ContinueAction_Type.TIMEOUT_STOP:
      return "TIMEOUT_STOP";
    case ContinueAction_Type.BOT_SUBSTITUTION:
      return "BOT_SUBSTITUTION";
    case ContinueAction_Type.NEXT_STAGE:
      return "NEXT_STAGE";
    case ContinueAction_Type.UNRECOGNIZED:
    default:
      return "UNRECOGNIZED";
  }
}

export enum ContinueAction_State {
  STATE_UNKNOWN = "STATE_UNKNOWN",
  BLOCKED = "BLOCKED",
  WAITING = "WAITING",
  READY_AUTO = "READY_AUTO",
  READY_MANUAL = "READY_MANUAL",
  UNRECOGNIZED = "UNRECOGNIZED",
}

export function continueAction_StateFromJSON(object: any): ContinueAction_State {
  switch (object) {
    case 0:
    case "STATE_UNKNOWN":
      return ContinueAction_State.STATE_UNKNOWN;
    case 1:
    case "BLOCKED":
      return ContinueAction_State.BLOCKED;
    case 2:
    case "WAITING":
      return ContinueAction_State.WAITING;
    case 3:
    case "READY_AUTO":
      return ContinueAction_State.READY_AUTO;
    case 4:
    case "READY_MANUAL":
      return ContinueAction_State.READY_MANUAL;
    case -1:
    case "UNRECOGNIZED":
    default:
      return ContinueAction_State.UNRECOGNIZED;
  }
}

export function continueAction_StateToJSON(object: ContinueAction_State): string {
  switch (object) {
    case ContinueAction_State.STATE_UNKNOWN:
      return "STATE_UNKNOWN";
    case ContinueAction_State.BLOCKED:
      return "BLOCKED";
    case ContinueAction_State.WAITING:
      return "WAITING";
    case ContinueAction_State.READY_AUTO:
      return "READY_AUTO";
    case ContinueAction_State.READY_MANUAL:
      return "READY_MANUAL";
    case ContinueAction_State.UNRECOGNIZED:
    default:
      return "UNRECOGNIZED";
  }
}

function createBaseGcState(): GcState {
  return { teamState: {}, autoRefState: {}, trackers: {}, continueActions: [] };
}

export const GcState = {
  fromJSON(object: any): GcState {
    return {
      teamState: isObject(object.teamState)
        ? Object.entries(object.teamState).reduce<{ [key: string]: GcStateTeam }>((acc, [key, value]) => {
          acc[key] = GcStateTeam.fromJSON(value);
          return acc;
        }, {})
        : {},
      autoRefState: isObject(object.autoRefState)
        ? Object.entries(object.autoRefState).reduce<{ [key: string]: GcStateAutoRef }>((acc, [key, value]) => {
          acc[key] = GcStateAutoRef.fromJSON(value);
          return acc;
        }, {})
        : {},
      trackers: isObject(object.trackers)
        ? Object.entries(object.trackers).reduce<{ [key: string]: string }>((acc, [key, value]) => {
          acc[key] = String(value);
          return acc;
        }, {})
        : {},
      continueActions: Array.isArray(object?.continueActions)
        ? object.continueActions.map((e: any) => ContinueAction.fromJSON(e))
        : [],
    };
  },

  toJSON(message: GcState): unknown {
    const obj: any = {};
    obj.teamState = {};
    if (message.teamState) {
      Object.entries(message.teamState).forEach(([k, v]) => {
        obj.teamState[k] = GcStateTeam.toJSON(v);
      });
    }
    obj.autoRefState = {};
    if (message.autoRefState) {
      Object.entries(message.autoRefState).forEach(([k, v]) => {
        obj.autoRefState[k] = GcStateAutoRef.toJSON(v);
      });
    }
    obj.trackers = {};
    if (message.trackers) {
      Object.entries(message.trackers).forEach(([k, v]) => {
        obj.trackers[k] = v;
      });
    }
    if (message.continueActions) {
      obj.continueActions = message.continueActions.map((e) => e ? ContinueAction.toJSON(e) : undefined);
    } else {
      obj.continueActions = [];
    }
    return obj;
  },
};

function createBaseGcState_TeamStateEntry(): GcState_TeamStateEntry {
  return { key: "", value: undefined };
}

export const GcState_TeamStateEntry = {
  fromJSON(object: any): GcState_TeamStateEntry {
    return {
      key: isSet(object.key) ? String(object.key) : "",
      value: isSet(object.value) ? GcStateTeam.fromJSON(object.value) : undefined,
    };
  },

  toJSON(message: GcState_TeamStateEntry): unknown {
    const obj: any = {};
    message.key !== undefined && (obj.key = message.key);
    message.value !== undefined && (obj.value = message.value ? GcStateTeam.toJSON(message.value) : undefined);
    return obj;
  },
};

function createBaseGcState_AutoRefStateEntry(): GcState_AutoRefStateEntry {
  return { key: "", value: undefined };
}

export const GcState_AutoRefStateEntry = {
  fromJSON(object: any): GcState_AutoRefStateEntry {
    return {
      key: isSet(object.key) ? String(object.key) : "",
      value: isSet(object.value) ? GcStateAutoRef.fromJSON(object.value) : undefined,
    };
  },

  toJSON(message: GcState_AutoRefStateEntry): unknown {
    const obj: any = {};
    message.key !== undefined && (obj.key = message.key);
    message.value !== undefined && (obj.value = message.value ? GcStateAutoRef.toJSON(message.value) : undefined);
    return obj;
  },
};

function createBaseGcState_TrackersEntry(): GcState_TrackersEntry {
  return { key: "", value: "" };
}

export const GcState_TrackersEntry = {
  fromJSON(object: any): GcState_TrackersEntry {
    return { key: isSet(object.key) ? String(object.key) : "", value: isSet(object.value) ? String(object.value) : "" };
  },

  toJSON(message: GcState_TrackersEntry): unknown {
    const obj: any = {};
    message.key !== undefined && (obj.key = message.key);
    message.value !== undefined && (obj.value = message.value);
    return obj;
  },
};

function createBaseGcStateTeam(): GcStateTeam {
  return {
    connected: false,
    connectionVerified: false,
    remoteControlConnected: false,
    remoteControlConnectionVerified: false,
    advantageChoice: undefined,
  };
}

export const GcStateTeam = {
  fromJSON(object: any): GcStateTeam {
    return {
      connected: isSet(object.connected) ? Boolean(object.connected) : false,
      connectionVerified: isSet(object.connectionVerified) ? Boolean(object.connectionVerified) : false,
      remoteControlConnected: isSet(object.remoteControlConnected) ? Boolean(object.remoteControlConnected) : false,
      remoteControlConnectionVerified: isSet(object.remoteControlConnectionVerified)
        ? Boolean(object.remoteControlConnectionVerified)
        : false,
      advantageChoice: isSet(object.advantageChoice) ? TeamAdvantageChoice.fromJSON(object.advantageChoice) : undefined,
    };
  },

  toJSON(message: GcStateTeam): unknown {
    const obj: any = {};
    message.connected !== undefined && (obj.connected = message.connected);
    message.connectionVerified !== undefined && (obj.connectionVerified = message.connectionVerified);
    message.remoteControlConnected !== undefined && (obj.remoteControlConnected = message.remoteControlConnected);
    message.remoteControlConnectionVerified !== undefined &&
      (obj.remoteControlConnectionVerified = message.remoteControlConnectionVerified);
    message.advantageChoice !== undefined &&
      (obj.advantageChoice = message.advantageChoice ? TeamAdvantageChoice.toJSON(message.advantageChoice) : undefined);
    return obj;
  },
};

function createBaseTeamAdvantageChoice(): TeamAdvantageChoice {
  return { choice: TeamAdvantageChoice_AdvantageChoice.STOP };
}

export const TeamAdvantageChoice = {
  fromJSON(object: any): TeamAdvantageChoice {
    return {
      choice: isSet(object.choice)
        ? teamAdvantageChoice_AdvantageChoiceFromJSON(object.choice)
        : TeamAdvantageChoice_AdvantageChoice.STOP,
    };
  },

  toJSON(message: TeamAdvantageChoice): unknown {
    const obj: any = {};
    message.choice !== undefined && (obj.choice = teamAdvantageChoice_AdvantageChoiceToJSON(message.choice));
    return obj;
  },
};

function createBaseGcStateAutoRef(): GcStateAutoRef {
  return { connectionVerified: false };
}

export const GcStateAutoRef = {
  fromJSON(object: any): GcStateAutoRef {
    return { connectionVerified: isSet(object.connectionVerified) ? Boolean(object.connectionVerified) : false };
  },

  toJSON(message: GcStateAutoRef): unknown {
    const obj: any = {};
    message.connectionVerified !== undefined && (obj.connectionVerified = message.connectionVerified);
    return obj;
  },
};

function createBaseGcStateTracker(): GcStateTracker {
  return { sourceName: "", uuid: "", ball: undefined, robots: [] };
}

export const GcStateTracker = {
  fromJSON(object: any): GcStateTracker {
    return {
      sourceName: isSet(object.sourceName) ? String(object.sourceName) : "",
      uuid: isSet(object.uuid) ? String(object.uuid) : "",
      ball: isSet(object.ball) ? Ball.fromJSON(object.ball) : undefined,
      robots: Array.isArray(object?.robots) ? object.robots.map((e: any) => Robot.fromJSON(e)) : [],
    };
  },

  toJSON(message: GcStateTracker): unknown {
    const obj: any = {};
    message.sourceName !== undefined && (obj.sourceName = message.sourceName);
    message.uuid !== undefined && (obj.uuid = message.uuid);
    message.ball !== undefined && (obj.ball = message.ball ? Ball.toJSON(message.ball) : undefined);
    if (message.robots) {
      obj.robots = message.robots.map((e) => e ? Robot.toJSON(e) : undefined);
    } else {
      obj.robots = [];
    }
    return obj;
  },
};

function createBaseBall(): Ball {
  return { pos: undefined, vel: undefined };
}

export const Ball = {
  fromJSON(object: any): Ball {
    return {
      pos: isSet(object.pos) ? Vector3.fromJSON(object.pos) : undefined,
      vel: isSet(object.vel) ? Vector3.fromJSON(object.vel) : undefined,
    };
  },

  toJSON(message: Ball): unknown {
    const obj: any = {};
    message.pos !== undefined && (obj.pos = message.pos ? Vector3.toJSON(message.pos) : undefined);
    message.vel !== undefined && (obj.vel = message.vel ? Vector3.toJSON(message.vel) : undefined);
    return obj;
  },
};

function createBaseRobot(): Robot {
  return { id: undefined, pos: undefined };
}

export const Robot = {
  fromJSON(object: any): Robot {
    return {
      id: isSet(object.id) ? RobotId.fromJSON(object.id) : undefined,
      pos: isSet(object.pos) ? Vector2.fromJSON(object.pos) : undefined,
    };
  },

  toJSON(message: Robot): unknown {
    const obj: any = {};
    message.id !== undefined && (obj.id = message.id ? RobotId.toJSON(message.id) : undefined);
    message.pos !== undefined && (obj.pos = message.pos ? Vector2.toJSON(message.pos) : undefined);
    return obj;
  },
};

function createBaseContinueAction(): ContinueAction {
  return {
    type: ContinueAction_Type.TYPE_UNKNOWN,
    forTeam: Team.UNKNOWN,
    continuationIssues: [],
    readyAt: undefined,
    state: ContinueAction_State.STATE_UNKNOWN,
  };
}

export const ContinueAction = {
  fromJSON(object: any): ContinueAction {
    return {
      type: isSet(object.type) ? continueAction_TypeFromJSON(object.type) : ContinueAction_Type.TYPE_UNKNOWN,
      forTeam: isSet(object.forTeam) ? teamFromJSON(object.forTeam) : Team.UNKNOWN,
      continuationIssues: Array.isArray(object?.continuationIssues)
        ? object.continuationIssues.map((e: any) => String(e))
        : [],
      readyAt: isSet(object.readyAt) ? fromJsonTimestamp(object.readyAt) : undefined,
      state: isSet(object.state) ? continueAction_StateFromJSON(object.state) : ContinueAction_State.STATE_UNKNOWN,
    };
  },

  toJSON(message: ContinueAction): unknown {
    const obj: any = {};
    message.type !== undefined && (obj.type = continueAction_TypeToJSON(message.type));
    message.forTeam !== undefined && (obj.forTeam = teamToJSON(message.forTeam));
    if (message.continuationIssues) {
      obj.continuationIssues = message.continuationIssues.map((e) => e);
    } else {
      obj.continuationIssues = [];
    }
    message.readyAt !== undefined && (obj.readyAt = message.readyAt.toISOString());
    message.state !== undefined && (obj.state = continueAction_StateToJSON(message.state));
    return obj;
  },
};

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

function isObject(value: any): boolean {
  return typeof value === "object" && value !== null;
}

function isSet(value: any): boolean {
  return value !== null && value !== undefined;
}
