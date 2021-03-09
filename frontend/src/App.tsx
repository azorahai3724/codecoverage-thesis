import React from 'react';
import logo from './logo.svg';
import './App.css';
import { useForm } from "react-hook-form"

type NewApp = {
  Name: string;
  CoverageFile: FileList;
};


//https://www.carlrippon.com/getting-started-with-react-hook-form-with-typeScript/

export const AppForm = () => {
  const { register, handleSubmit, errors } = useForm<NewApp>();
  const fd = new FormData();
  const onNewAppSubmit = async (d: NewApp) => {
    fd.append('Name', d.Name)
    fd.append('CoverageFile', d.CoverageFile[0])
    const response = await fetch("http://localhost:10000/app/create", {
      method: 'POST',
      mode:'cors',
      body: fd,
      headers : { 
        'Accept': 'application/json'
       }
    });
    //const jsonData = await response.json();
    //console.log("Response: ", jsonData);
    //console.log(response.text)
  };

  return (
    <form onSubmit={handleSubmit(onNewAppSubmit)}>
      <div className="field">
        <label htmlFor="Name">Name:</label>
        <input type="text" id="Name" name="Name" ref={register({required:true})}/>
        {errors.Name && errors.Name.type === "required" && (
          <div className="error">App name is mandatory.</div>
        )}
      </div>
      <div className="field">
        <label htmlFor="CoverageFile">Coverage File: </label>
        <input type="file" id="CoverageFile" name="CoverageFile" ref={register({required:true})}/>
        {errors.CoverageFile && (
          <div className="error">Coverage file is mandatory.</div>
        )}
      </div>
      <button type="submit">Create app</button>
    </form>

  );

};

