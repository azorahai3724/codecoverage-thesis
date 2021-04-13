import React from 'react';
import { render, screen } from '@testing-library/react';
import {AppForm} from './App';

test('renders app name label ', () => {
  render(<AppForm />);
  const formElement = screen.getByText(/Name:/i);
  expect(formElement).toBeInTheDocument();
});

test('renders coverage file button label ', () => {
  render(<AppForm />);
  const formElement = screen.getByText(/Coverage File:/i);
  expect(formElement).toBeInTheDocument();
});
