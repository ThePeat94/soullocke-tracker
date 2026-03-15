<script lang="ts">
	import type { InputType } from '$components/textinput/types';

	type Props = {
		value: string
		label: string
		placeholder?: string
		type?: InputType
	};

	let {
		value = $bindable<string>(),
		label,
		placeholder,
		type = 'text',
	} : Props = $props();
</script>


<label class="label">
	<span class="label-text">{label}</span>
	<input type={type} class="input p-3 rounded-md border hover:border-gray-50" bind:value={value} placeholder={placeholder}/>
</label>
