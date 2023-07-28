<script lang="ts">
	import { onMount } from 'svelte';
	import type { PageData } from './$types';
	export let data: PageData;

	let userID: string | null = null;
	let userName: string | null = null;

	let ws: WebSocket;

	import { messages, users, currentUser } from '$lib/messagesStores';

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
			$currentUser = userID;

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
				$users[item.user_id] = item.user_name;
			});

			$messages = [...$messages, prevMessages][0];

			// websockets
			ws = new WebSocket('ws://localhost:8081/ws?user_id=' + userID + '&chat_id=' + data.chat.id);
			ws.onmessage = async (event) => {
				// @ts-ignore
				const wsMessage = JSON.parse(event.data);
				$messages = [...$messages, wsMessage];
				// @ts-ignore
				if ($users[wsMessage.user_id] === undefined) {
					console.log('new user detected');
					const user = await fetch(`http://localhost:8081/users/${wsMessage.user_id}`, {
						method: 'GET'
					}).then((res) => res.json());
					// @ts-ignore
					$users[user.user_id] = user.name;
					$messages = $messages;
					console.log($users);
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
</script>

<!-- svelte-ignore non-top-level-reactive-declaration -->
{#if data.exists}
	Chat exists with id {data.chat.id} and name '{data.chat.name}' with description '{data.chat
		.description}' You are currently logged in as '{userName}' with id '{userID}'

	{#each $messages as message}
		<Message {message} />
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
