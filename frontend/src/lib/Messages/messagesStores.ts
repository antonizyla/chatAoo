import { writable } from "svelte/store";

export type message = {
    message_id: string
    user_id: string
    body: string
    created_at: string
    updated_at: string
    deleted: boolean
    chat_id: string
}

export interface reaction {
    user_id: string
    emoji: string
}

export interface messageReaction extends reaction {
    message_id: string
}

let msgs: message[] = []
export let messages = writable(msgs)

export let users = writable(new Map<string, string>())
export let currentUser = writable("")

// holds a key of message to a list of reactions
export let reactions = writable(new Map<string, reaction[]>())

export let wsPayload = writable("Empty")
