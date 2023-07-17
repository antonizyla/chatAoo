<script lang="ts">
	import { onMount } from 'svelte';
	import type { PageData } from './$types';
	export let data: PageData;

	let userID: string | null = null;
	let userName: string | null = null;

	let ws: WebSocket;

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
			const apiQuery = `http://localhost:8081/getMessages?chat_id=${
				data.chat.id
			}&timestamp=${Date.now()}`;
			console.log(apiQuery);
			const prevMessages = await fetch(
				`http://localhost:8081/getMessages?chat_id=${data.chat.id}&timestamp=${Date.now()}`
			).then((res) => res.json());

			messages = [...messages, prevMessages][0];

			// websockets
			ws = new WebSocket('ws://localhost:8081/ws?userID=' + userID + '&chat_id=' + data.chat.id);
			ws.onmessage = (event) => {
				messages = [...messages, JSON.parse(event.data)];
			};
		}
	});

	type msg = {
		msg_id: string;
		content: string;
		sender_id: string;
		chat: string;
		sender_name: string;
		created_at: string;
	};
	let messages: msg[] = [];
	let msg: string = '';

	function sendMessage() {
		if (msg != '') {
			ws.send(
				JSON.stringify({
					content: msg,
					sender_id: userID,
					chat_id: data.chat.id
				})
			);
			msg = '';
		}
	}

	import Message from '$lib/Message.svelte';
</script>

<!-- svelte-ignore non-top-level-reactive-declaration -->
{#if data.exists}
	Chat exists with id {data.chat.id} and name '{data.chat.name}' with description '{data.chat
		.description}' You are currently logged in as '{userName}' with id '{userID}'

	{#each messages as message}
		<Message
			sender_name={message.sender_name}
			content={message.content}
			time={message.created_at}
			alignRight={message.sender_id == userID}
		/>
	{/each}

	<input
		type="text"
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
