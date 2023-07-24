<script lang="ts">
	import { onMount } from 'svelte';
	import type { PageData } from './$types';
	export let data: PageData;

	let userID: string | null = null;
	let userName: string | null = null;

	let ws: WebSocket;

	let messages: any[] = [];
	let users = new Map<string, string>();

	onMount(async () => {
		// check if user exists on browser
		// if not, redirect to login page

		userID = localStorage.getItem('userID');
		userName = localStorage.getItem('userName');
		if (!userID) {
			window.location.href = '/';
			// create store to store chat of user
			localStorage.setItem('chatId', data.chat.id);
		} else {
			// load in the chats for the user
			const prevMessages = await fetch(
				`http://localhost:8081/messages/${data.chat.id}/${Date.now()}`,
				{ method: 'GET' }
			).then((res) => res.json());

			// load the users that are currently in the chat
			const prevUsers = await fetch(`http://localhost:8081/chats/${data.chat.id}/users`, {
				method: 'GET'
			}).then((res) => res.json());

			prevUsers.map((item: { user_name: string; user_id: string }) => {
				// @ts-ignore
				users[item.user_id] = item.user_name;
			});

			messages = [...messages, prevMessages][0];

			// websockets
			ws = new WebSocket('ws://localhost:8081/ws?user_id=' + userID + '&chat_id=' + data.chat.id);
			ws.onmessage = async (event) => {
				// @ts-ignore
				messages = [...messages, JSON.parse(event.data)];
                console.log(event.data)
				if (!users.has(event.data.user_id)) {
					const user = await fetch(`http://localhost:8081/users/${event.data.user_id}`, {
						method: 'GET'
					}).then((res) => res.json());
					users.set(event.data.user_id, user.name);
				}
			};
		}
	});

	let msg: string = '';

	function sendMessage() {
		if (msg != '') {
			const payload = {
				action_type: 'newMessage',
				chat_id: data.chat.id,
				user_id: userID,
				message: msg
			};
			ws.send(JSON.stringify(payload));
			msg = '';
		}
	}

	import Message from '$lib/Message.svelte';
	import { json } from '@sveltejs/kit';
</script>

<!-- svelte-ignore non-top-level-reactive-declaration -->
{#if data.exists}
	Chat exists with id {data.chat.id} and name '{data.chat.name}' with description '{data.chat
		.description}' You are currently logged in as '{userName}' with id '{userID}'

	{#each messages as message}
		<Message
			sender_name={users[message.user_id]}
			body={message.body}
			time={message.created_at}
			alignRight={message.sender_id == userID}
		/>
	{/each}

	<input
		type="text"
		class="border-gray-400 border"
		name="msg"
		id="msg"
		bind:value={msg}
		on:keydown={(e) => {
			if (e.key === 'Enter' && msg != '') {
				sendMessage();
			}
		}}
	/>
	<button on:click={sendMessage}>SEND</button>
{:else}
	<h1>404</h1>
	<p>Chat does not exist</p>
{/if}
