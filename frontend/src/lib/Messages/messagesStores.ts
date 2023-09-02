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

export type reaction = {
    message_id: string
    user_id: string
    emoji: string
}

let msgs: message[] = []
export let messages = writable(msgs)

export let users = writable(new Map<string, string>())
export let currentUser = writable("")

export let reactions = writable(new Map<string, { user_id: string; emoji: string }[]>())

export let wsPayload = writable("Empty")
