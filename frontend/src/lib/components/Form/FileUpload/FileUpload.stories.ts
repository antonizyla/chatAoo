import type { Meta, StoryObj } from '@storybook/svelte';

import FileUpload from './FileUpload.svelte';

const meta = {
	title: 'Form/FileUpload',
	component: FileUpload,
	tags: ['autodocs'],
	argTypes: {
		name: { control: 'text' }
	}
} satisfies Meta<FileUpload>;

export default meta;
type Story = StoryObj<typeof meta>;

export const Primary: Story = {
	args: {
		name: 'file dasd'
	}
};
