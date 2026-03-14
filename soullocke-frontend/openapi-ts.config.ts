import { defineConfig } from '@hey-api/openapi-ts';

export default defineConfig({
	input: '../openapi/openapi.yaml',
	output: 'src/lib/api/generated',
	plugins: [
		'@hey-api/client-fetch',
		'@hey-api/typescript',
		{
			name: '@tanstack/svelte-query',
			queryOptions: true,
			mutationOptions: true,
			queryKeys: true
		}
	]
});
