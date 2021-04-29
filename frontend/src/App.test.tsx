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

test('renders commit hash label ', () => {
  render(<AppForm />);
  const formElement = screen.getByText(/Commit Hash/i);
  expect(formElement).toBeInTheDocument();
});

test('renders commit hash label ', () => {
  render(<AppForm />);
  const formElement = screen.getByText(/Commit Hash/i);
  expect(formElement).toBeInTheDocument();
});

test('renders form ', () => {
  const element = render(<AppForm />);
  expect(element.queryByTestId("/form/i")).toBeTruthy;
});

test('renders app name input ', () => {
  const element = render(<AppForm />);
  expect(element.queryByTestId("/name/i")).toBeTruthy;
});

test('renders file input ', () => {
  const element = render(<AppForm />);
  expect(element.queryByTestId("/file/i")).toBeTruthy;
});

test('renders commit hash input ', () => {
  const element = render(<AppForm />);
  expect(element.queryByTestId("/CommitText/i")).toBeTruthy;
});

test('renders submit button ', () => {
  const element = render(<AppForm />);
  expect(element.queryByTestId("/submit/i")).toBeTruthy;
});

test('renders submit button text ', () => {
  render(<AppForm />);
  const formElement = screen.getByText(/Upload/i);
  expect(formElement).toBeInTheDocument();
});

test('not renders app id if no input is given', () => {
  const element = render(<AppForm />);
  expect(element.queryByTestId("/appName/i")).toBeFalsy;
});

test('not renders app name if no input is given', () => {
  const element = render(<AppForm />);
  expect(element.queryByTestId("/appName/i")).toBeFalsy;
});

test('not renders graph if no input is given', () => {
  const element = render(<AppForm />);
  expect(element.queryByTestId("/graph/i")).toBeEmptyDOMElement;
});

test('not renders table if no input is given', () => {
  const element = render(<AppForm />);
  expect(element.queryByTestId("/table/i")).toBeEmptyDOMElement;
});
