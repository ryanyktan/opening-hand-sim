import React from 'react';
import { Form, Field, Formik, ErrorMessage } from "formik";
import '../styles/OpeningHandSim.css'

const OpeningHandSim = () => {

    return (
        <div>
            <h1>
                opening hand simulator
            </h1>
            <Formik
                initialValues={{
                    decklist: '',
                    numberOfSimulations: '',
                }}
                onSubmit={async (values) => {
                    await new Promise((r) => setTimeout(r, 500));
                    alert(JSON.stringify(values, null, 2));
                }}
            >
                <Form>
                    <label htmlFor="decklist">Decklist</label>
                    <Field id="decklist" name="decklist" placeholder="Enter your decklist here" as="textarea" />

                    <div>
                        <label htmlFor="numberOfSimulations">Number Of Simulations:</label>
                        <Field name="numberOfSimulations" as="select" label="Number of Simulations">
                            <option value="">Choose how many simulations to run</option>
                            <option value="10">10</option>
                            <option value="100">100</option>
                            <option value="1000">1000</option>
                        </Field>
                        <ErrorMessage name="numberOfSimulations" component="div" />
                    </div>

                    <button type="submit">Submit</button>
                </Form>
            </Formik>
        </div>
    );
}

export default OpeningHandSim
