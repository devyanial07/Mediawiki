package main

import(
		"bytes"
			"encoding/json"
				"fmt"
					"io/ioutil"
						"log"
							"net/http"
								"os"
							)

							type apiResponse struct {
									Id, Kind, LongUrl string
								}

								func main(){
										longUrl := os.Args[len(os.Args)-1]
											
											body := bytes.newBufferString(fmt.Spritf(
													`{"longUrl" : %s}`,
														longUrl))
															
															request, err := http.NewRequest("POST",
																"https://www.googleapis.com/urlshortner/v1/curl", 
																	body)
																		
																		request.Header.Add("Content-Type", "application/json")
																			
																			client := http.Client {}
																				
																				response, err := client.Do(request)
																					
																					if err != nil {
																								log.Fatal(err)
																									}
																										
																										outputBytes, err := ioutil.ReadAll(response.Body)
																											response.Body.Close()
																												
																												var output apiResponse
																													err = json.Unmarshal(outputBytes, &output)
																														
																														if err != nil {
																																	log.Fatal(err)
																																		}
																																			
																																			fmt.Printf("%s\n", output.Id)
																																		}
