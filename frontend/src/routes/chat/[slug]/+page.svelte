<script lang="ts">
	import { onMount } from 'svelte';
	import type { PageData } from './$types';
	export let data: PageData;

	let userID: string | null = null;
	let userName: string | null = null;

	let ws: WebSocket;

	import { messages, users, currentUser, wsPayload, reactions } from '$lib/Messages/messagesStores';

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
				`http://localhost:8081/chats/${data.chat.id}/messages/${Date.now()}`,
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

			// load the reactions of the messages
			const chatReactions = await fetch(`http://localhost:8081/chats/${data.chat.id}/reactions`, {
				method: 'GET'
			}).then((res) => res.json());

			chatReactions.map((item: any) => {
				const itemReaction = { user_id: item.user_id, reaction: item.reaction };
				if (!$reactions[item.message_id as keyof typeof $reactions]) {
					$reactions[item.message_id] = [itemReaction];
				} else {
					$reactions[item.message_id].push(itemReaction);
				}
			});

			$messages = [...$messages, prevMessages][0];

			// websockets
			ws = new WebSocket('ws://localhost:8081/ws?user_id=' + userID + '&chat_id=' + data.chat.id);
			ws.onmessage = async (event) => {
				// @ts-ignore
				const wsMessage = JSON.parse(event.data);
				console.log(wsMessage);
				if (wsMessage.user_id && wsMessage.chat_id && wsMessage.message_id) {
					// new message received
					console.log('New Message received');
					$messages = [...$messages, wsMessage];
					// @ts-ignore
					if ($users[wsMessage.user_id] === undefined) {
						console.log('new user detected');
						const user = await fetch(`http://localhost:8081/users/${wsMessage.user_id}`, {
							method: 'GET'
						}).then((res) => res.json());
						// @ts-ignore
						$users[user.user_id] = user.name;
					}
				} else if (wsMessage.message_id && wsMessage.updated_at) {
					console.log('Message Deleted detected');
					$messages.forEach((msg) => {
						if (wsMessage.message_id == msg.message_id) {
							msg.deleted = true;
							msg.updated_at = wsMessage.updated_at;
						}
					});
				} else if (
					wsMessage.message_id &&
					wsMessage.user_id &&
					wsMessage.reaction_emoji &&
					wsMessage.action_type == 'addReaction'
				) {
					const newReaction = { user_id: wsMessage.user_id, reaction: wsMessage.reaction_emoji };
					if (!$reactions[wsMessage.message_id]) {
						$reactions[wsMessage.message_id] = [newReaction];
					} else {
						$reactions[wsMessage.message_id].push(newReaction);
					}
					$reactions = $reactions;
				} else if (
					wsMessage.reaction_emoji &&
					wsMessage.user_id &&
					wsMessage.message_id &&
					wsMessage.action_type == 'removeReaction'
				) {
					console.log('remove reaction ws recieved');
					for (let i = 0; i < $reactions[wsMessage.message_id].length; i++) {
						if (
							$reactions[wsMessage.message_id][i].user_id === wsMessage.user_id &&
							$reactions[wsMessage.message_id][i].reaction === wsMessage.reaction_emoji
						) {
							$reactions[wsMessage.message_id].splice(i, 1);
						}
					}
					$reactions = $reactions;
				}
				$messages = $messages;
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
			$wsPayload = JSON.stringify(payload);
			msg = '';
		}
	}

	// everytime $payload changes send it to WebSocket
	function sendws(_payload: any) {
		if (ws) {
			ws.send($wsPayload);
		}
	}

	$: sendws($wsPayload);

	import Message from '$lib/Messages/Message.svelte';
	import ListUsers from '$lib/Users/ListUsers.svelte';
	import Button from '$lib/components/Button/Button.svelte';
</script>

{#if data.exists}
	<a href="/chat"><Button>Navigate To Your Chats</Button></a>
	Chat exists with id {data.chat.id} and name '{data.chat.name}' with description '{data.chat
		.description}' You are currently logged in as '{userName}' with id '{userID}'

	<ListUsers chat={data.chat.id} />

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
	<a href="/chat"><Button>Navigate To Your Chats</Button></a>
{/if}
