<script lang="ts">
	import Card from '$components/Card.svelte';
	import PrimaryButton from '$components/buttons/PrimaryButton.svelte';
	import TextInput from '$components/textinput/TextInput.svelte';
	import { getLobbyCreationMutation, getLobbyQuery } from '$api/lobby';
	import Combobox from '$components/Combobox.svelte';

	let lobbyId = $state('');
	const lobbyQuery = getLobbyQuery(() => lobbyId);

	let lobbyName = $state('');
	let lobbyPassword = $state('')
	let gameEdition = $state<string>();

	const createLobbyMutation = getLobbyCreationMutation();

	const handleCreateLobbyClick = (): void => {
		if (!gameEdition) {
			return;
		}
		createLobbyMutation.mutate({
			body: {
				name: lobbyName,
				password: lobbyPassword,
				gameEditionId: gameEdition,
			}
		});
	};


</script>

<h1 class="text-center text-3xl">
	Lobby Management
</h1>



<div class="mx-auto flex max-w-1/2 flex-col gap-4 mt-10">
	<Card>
		{#snippet header()}
			<h2 class="h4">Create Lobby</h2>
		{/snippet}
		{#snippet content()}
			<div class="flex flex-col gap-4">
				<TextInput label="Lobby Name" bind:value={lobbyName} disabled={createLobbyMutation.isPending} />
				<TextInput label="Password" type="password" bind:value={lobbyPassword} disabled={createLobbyMutation.isPending} />
				<Combobox
					items={[
						{ label: 'FireRed', value: 'firered', group: 'Generation 3 - Remake' },
						{ label: 'LeafGreen', value: 'leafgreen', group: 'Generation 3 - Remake' },
						{ label: 'Emerald', value: 'emerald', group: 'Generation 3' },
						{ label: 'Ruby', value: 'ruby', group: 'Generation 3' },
						{ label: 'Sapphire', value: 'sapphire', group: 'Generation 3' },
					]}
					label="Game Edition"
					bind:value={gameEdition}
				/>
			</div>
		{/snippet}
		{#snippet footer()}
			<PrimaryButton variant="filled" onClick={handleCreateLobbyClick} disabled={createLobbyMutation.isPending}>Create Lobby</PrimaryButton>
		{/snippet}
	</Card>
	<Card>
		{#snippet header()}
			<h2 class="h4">Get Lobby</h2>
		{/snippet}
		{#snippet content()}
			<TextInput label="Lobby Name" bind:value={lobbyId} />
			{#if lobbyQuery.isSuccess}
				{lobbyQuery.data.name}
			{/if}
			{#if lobbyQuery.isError}
				<p class="text-red-500">Error fetching lobby</p>
			{/if}
		{/snippet}
	</Card>
	{#if createLobbyMutation.isSuccess && createLobbyMutation.data}
		<div class="col-start-2">
			<Card>
				{#snippet header()}
					<h2 class="h4">Lobby Created</h2>
				{/snippet}
				{#snippet content()}
					<p>Lobby ID: {createLobbyMutation.data.lobbyId}</p>
				{/snippet}
			</Card>
		</div>
	{/if}
	{#if createLobbyMutation.isError}
		<div class="col-start-2">
			<Card>
				{#snippet header()}
					<h2 class="h4">Error Creating Lobby</h2>
				{/snippet}
				{#snippet content()}
					<p class="text-red-500">Error Error Error :(</p>
				{/snippet}
			</Card>
		</div>
	{/if}
</div>
