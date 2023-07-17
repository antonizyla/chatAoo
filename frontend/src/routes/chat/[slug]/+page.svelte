<script lang="ts">
	import { onMount } from 'svelte';
	import type { PageData } from './$types';
	export let data: PageData;

	let userID: string | null = null;
	let userName: string | null = null;

	let msg: string = '';

	let ws: WebSocket;

	type msg = {
		msgId: string;
		content: string;
		sender: string;
		chatID: string;
	};

	onMount(() => {
		// check if user exists on browser
		// if not, redirect to login page

		userID = localStorage.getItem('userID');
		userName = localStorage.getItem('userName');
		console.log(userID);
		if (!userID) {
			window.location.href = '/';
			// create store to store chat of user
			localStorage.setItem('chatId', data.chat.id);
		} else {
			// websockets
			ws = new WebSocket('ws://localhost:8081/ws?userID=' + userID + '&chat_id=' + data.chat.id);
			ws.onmessage = (event) => {
				messages = [...messages, JSON.parse(event.data)];
			};
		}
	});

	let messages: msg[] = [];
</script>

<!-- svelte-ignore non-top-level-reactive-declaration -->
{#if data.exists}
	Chat exists with id {data.chat.id} and name '{data.chat.name}' with description '{data.chat
		.description}' You are currently logged in as '{userName}' with id '{userID}'

	{#each messages as message}
		<p>{message.content} from {message.sender}</p>
	{/each}

	<input
		type="text"
		name="msg"
		id="msg"
		bind:value={msg}
		on:keydown={(e) => {
			if (e.keyCode === 13 && msg != '') {
				let message = {
					msgId: '',
					content: msg,
					sender: userID,
					chatID: data.chat.id
				};
				ws.send(JSON.stringify(message));
				msg = '';
			}
		}}
	/>
{:else}
	<h1>404</h1>
	<p>Chat does not exist</p>
{/if}
