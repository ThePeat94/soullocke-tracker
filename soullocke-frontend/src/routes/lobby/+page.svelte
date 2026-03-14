<script lang="ts">
	import Card from '../../components/Card.svelte';
	import PrimaryButton from '../../components/buttons/PrimaryButton.svelte';
	import TextInput from '../../components/TextInput.svelte';

	let lobbyName = $state('Test Lobby');


	const handleBtnClick = (): void => {
		console.log(lobbyName);
	};

</script>

<h1 class="text-center text-3xl">
	Lobby Management
</h1>



<div class="grid grid-cols-3 mt-10 gap-4">
	<div class="col-start-2">
		<Card>
			<header slot="header">
				<h2 class="h4">Create Lobby</h2>
			</header>
			<section slot="content">
				<TextInput label="Lobby Name" bind:value={lobbyName} />
			</section>
			<section slot="footer">
				<PrimaryButton variant="filled" onClick={handleBtnClick}>Create Lobby</PrimaryButton>
			</section>
		</Card>
	</div>
</div>
