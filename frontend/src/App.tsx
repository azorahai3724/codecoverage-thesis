import React from 'react';
import './App.css';
import { useForm } from "react-hook-form"
import ReactDOM from 'react-dom';
import { XAxis, YAxis, HorizontalGridLines, VerticalGridLines,LineMarkSeries, Hint, LineMarkSeriesPoint, FlexibleXYPlot, XYPlot, MarkSeries } from 'react-vis';
import "react-vis/dist/style.css";
import {Button, Table, TableBody, TableCell, TableHeader, TableHeaderCell, TableHeaderRow, TableRow } from '@bitrise/bitkit';

type NewApp = {
  Name: string;
  CoverageFile: FileList;
  CommitHash: string;
};


//https://www.carlrippon.com/getting-started-with-react-hook-form-with-typeScript/

export const AppForm = () => {

  const { register, handleSubmit, errors } = useForm<NewApp>();

  const fd = new FormData();

  const onNewAppSubmit = async (d: NewApp) => {

    fd.append('Name', d.Name)
    fd.append('CoverageFile', d.CoverageFile[0])
    fd.append('CommitHash', d.CommitHash)

    const response = await fetch("http://localhost:10000/app/create", {

      method: 'POST',
      mode:'cors',
      body: fd,
      headers : { 
        'Accept': 'application/json'
       }

    });

    const appJsonData = await response.json();
    
    var appName = appJsonData["Name"];
    var reports = appJsonData["Reports"];

    type graphData = {
      x: Date,
      y: number,
      z: string
    }

    let data: graphData[] = [];

    for (var i=0; i<reports.length; i++) {

        let golangCD = reports[i].CreationDate;
        let cdRaw = golangCD.slice(0,19);
        console.log(cdRaw)
        let cd = UTCToLocalTime(new Date(cdRaw));
        console.log("anan", cd)

        let cpString = reports[i].CoveragePercentage;
        let cp = Number(cpString);
        
        let ch = reports[i].CommitHash;
        data[i] = {x: cd, y: cp, z: ch};
        console.log(data[i].x)

    }

    class Graph extends React.Component {
      hoveredCircle!: LineMarkSeriesPoint;
      state = {
        hovered: false
      }
      render() {
        const {hovered} = this.state;
        return (
          <React.Fragment>

            <h2>{ appName }</h2>
            
            <XYPlot xType="time" yDomain={[0,100]} width={1300} height={900}>
            <VerticalGridLines />
            <HorizontalGridLines />
            <XAxis title="Date" />
            <YAxis title="Coverage percent" />
            <LineMarkSeries
              data={ data }
            />
            <MarkSeries
              data={ data }
              onValueMouseOver={d => {
                this.hoveredCircle = d;
                this.setState({hovered: true});
              }}
              onValueMouseOut={_d => this.setState({hovered: false})}            
            />
            { hovered === true && 
                <Hint value={this.hoveredCircle}>
                    {this.hoveredCircle.z}
                </Hint>
            }
            </XYPlot>

            <Table>
              <TableHeader>
                <TableHeaderRow>
                  <TableHeaderCell>Creation Date</TableHeaderCell>
                  <TableHeaderCell>Coverage percentage</TableHeaderCell>
                  <TableHeaderCell>Commit Hash</TableHeaderCell>
                </TableHeaderRow>
              </TableHeader>
              <TableBody>
              {data.map(d => 
                <TableRow>
                  <TableCell>{d.x.toString()}</TableCell>
                  <TableCell>{d.y}</TableCell>
                  <TableCell>{d.z}</TableCell>
                </TableRow>
                )
              }
              </TableBody>
          </Table>

          </React.Fragment>

        )
      }
    }

    const graphDiv = document.getElementById("graph");

    if (Object.keys(appJsonData).length !== 0) {
        ReactDOM.render(
          <Graph />, graphDiv
        )
    }
    
  };

  return (
      <form onSubmit={handleSubmit(onNewAppSubmit)}>
        <div className="field">
          <label htmlFor="Name">Name: </label>
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

        <div className="field">
          <label htmlFor="CommitHash">Commit Hash </label>
          <input type="text" id="CommitHash" name="CommitHash" ref={register({required:true})}/>
          {errors.CommitHash && (
            <div className="error">Commit Hash is mandatory.</div>
          )}
        </div>

        <Button size="small" type="submit">Upload</Button>

      </form>
  );

};

function UTCToLocalTime(date: Date) {

  let localTime = new Date(date.getTime() + date.getTimezoneOffset() * 60 * 1000);

  localTime.setHours(date.getHours() - date.getTimezoneOffset() / 60);

  return localTime;   
}

