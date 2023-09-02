import type { Meta, StoryObj } from '@storybook/svelte';

import Author from './Author.svelte';

const meta = {
	title: 'Blog/Author',
	component: Author,
	tags: ['autodocs'],
	argTypes: {
		name: { control: 'text' },
		email: { control: 'text' },
		avatar: { control: 'text' },
		bio: { control: 'text' }
	}
} satisfies Meta<Author>;

export default meta;
type Story = StoryObj<typeof meta>;

import image from './sampleImage2.jpg?w=200&h=200&picture';

export const Primary: Story = {
	args: {
		name: 'John Doe',
		email: 'john.doe@gmail.com',
		avatar: image,
		bio: 'Lorem ipsum dolor sit amet, qui minim labore adipisicing minim sint cillum sint consectetur cupidatat.'
	}
};
